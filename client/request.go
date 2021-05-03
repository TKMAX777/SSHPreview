package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type RequestHandler struct {
	client   *http.Client
	localAdd string
}

type HTTPPreviewData struct {
	Data     io.Reader
	FileName string
}

type HTTPResponseFormat struct {
	Status string
}

func NewRequestHandler(addr string) *RequestHandler {
	return &RequestHandler{new(http.Client), addr}
}

func (r RequestHandler) ChromeShow() (err error) {
	_, err = http.Get(r.localAdd + "/chromeshow")

	return

}

func (r RequestHandler) ChromeOff() (err error) {
	_, err = http.Get(r.localAdd + "/chromeoff")
	return
}

func (r *RequestHandler) Send(data HTTPPreviewData) (err error) {
	var body = new(bytes.Buffer)
	var mw = multipart.NewWriter(body)

	fw, err := mw.CreateFormFile("file", data.FileName)
	if err != nil {
		return
	}

	io.Copy(fw, data.Data)
	var contentType = mw.FormDataContentType()

	mw.Close()

	resp, err := http.Post(r.localAdd+"/remote", contentType, body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var status HTTPResponseFormat
	json.NewDecoder(resp.Body).Decode(&status)

	if status.Status != "OK" {
		fmt.Println(status.Status)
	}

	return
}
