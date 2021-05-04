package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type HTTPRemoteRequest struct {
	preview chan HTTPPreviewData
	message chan string
}

type HTTPPreviewData struct {
	Data     io.Reader
	FileName string
}

func NewRemoteRequestHandler(preview chan HTTPPreviewData, message chan string) *HTTPRemoteRequest {
	return &HTTPRemoteRequest{preview, message}
}

func (h *HTTPRemoteRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.RemoteAddr, "127.0.0.1") {
		SendJSON(w, HTTPResponseFormat{Status: "Forbidden"}, 403)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		SendJSON(w, HTTPResponseFormat{Status: err.Error()}, 500)
		return
	}
	defer file.Close()

	if GetContentType(filepath.Ext(header.Filename)) == "" {
		SendJSON(w, HTTPResponseFormat{Status: "Invalid type"}, 400)
		return
	}

	var data = new(bytes.Buffer)
	io.Copy(data, file)

	go func() {
		b, _ := json.Marshal(HTTPResponseFormat{Status: "OK", Content: struct{ FileName string }{header.Filename}})
		h.message <- string(b)
	}()

	go func() {
		h.preview <- HTTPPreviewData{
			Data:     data,
			FileName: header.Filename,
		}
	}()

	SendJSON(w, HTTPResponseFormat{Status: "OK"}, 200)
	return
}

func (h *HTTPRemoteRequest) VerifyName(w http.ResponseWriter, r *http.Request) {
	if GetContentType(filepath.Ext(r.URL.Query().Get("name"))) == "" {
		SendJSON(w, HTTPResponseFormat{Status: "Not Implemented"}, 501)
		return
	}
	SendJSON(w, HTTPResponseFormat{Status: "OK"}, 200)
	return
}
