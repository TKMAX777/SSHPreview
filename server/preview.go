package main

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type HTTPPreviewSend struct {
	preview chan HTTPPreviewData
}

func NewPreviewSender(preview chan HTTPPreviewData) *HTTPPreviewSend {
	return &HTTPPreviewSend{preview}
}

func (h *HTTPPreviewSend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.RemoteAddr, "127.0.0.1") {
		w.WriteHeader(403)
		fmt.Fprintf(w, "Forbidden")
		return
	}

	var preview = <-h.preview
	var extension = filepath.Ext(preview.FileName)
	var contentType = GetContentType(extension)

	if contentType == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Illigal extension:%s", extension)
		return
	}

	w.Header().Add("Content-Type", contentType)
	io.Copy(w, preview.Data)

	return
}

func GetContentType(extension string) string {
	switch extension {
	case ".jpg":
		return "image/jpeg"
	case ".jpeg", ".gif", ".png":
		return "image/" + extension[1:]
	case ".mp3", ".wav", ".m4a", ".aac":
		return "audio/" + extension[1:]
	}
	return ""
}
