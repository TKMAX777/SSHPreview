package main

import (
	"net/http"
)

type HTTPHandler struct {
}

func NewHTTPHandler() *HTTPHandler {
	var h HTTPHandler
	var preview = make(chan HTTPPreviewData)
	var message = make(chan string)

	var chrome = NewHTTPChromeHandler()

	var remoteRequests = NewRemoteRequestHandler(preview, message)

	http.Handle("/", http.FileServer(http.Dir("resources")))

	http.HandleFunc("/ping",
		func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte("OK")) },
	)

	http.Handle("/remote", remoteRequests)
	http.HandleFunc("/verify", remoteRequests.VerifyName)

	http.Handle("/local", NewFileNameSender(message))
	http.Handle("/file", NewPreviewSender(preview))

	http.HandleFunc("/chromeshow", chrome.Show)
	http.HandleFunc("/chromeoff", chrome.Off)

	return &h
}

func (h *HTTPHandler) Start() (err error) {
	if err := http.ListenAndServe("localhost:"+Settings.ListenPort, nil); err != nil {
		return err
	}

	return
}
