package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
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

const (
	SockTypeTCP = iota
	SockTypeUNIX
)

func NewRequestHandler(addr string, sockType int) *RequestHandler {

	var client *http.Client
	switch sockType {
	case SockTypeUNIX:
		var sock = filepath.Join(addr, "http.sock")
		client = &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", sock)
				},
			},
		}
		addr = "."
	case SockTypeTCP:
		client = new(http.Client)
	}

	return &RequestHandler{client, "http://" + addr}
}

func (r RequestHandler) ChromeShow() (err error) {
	_, err = r.client.Get(r.localAdd + "/chromeshow")

	return

}

func (r RequestHandler) ChromeOff() (err error) {
	_, err = r.client.Get(r.localAdd + "/chromeoff")
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

	resp, err := r.client.Post(r.localAdd+"/remote", contentType, body)
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

func (r *RequestHandler) Verify(fileName string) (err error) {
	resp, err := r.client.Get(r.localAdd + "/verify?name=" + url.QueryEscape(fileName))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Invalid")
	}
	return
}

func (r *RequestHandler) Ping() (err error) {
	_, err = r.client.Get(r.localAdd + "/ping")
	return err
}
