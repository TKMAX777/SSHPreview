package main

import (
	"fmt"
	"net/http"
	"strings"
)

type HTTPChromeSwitch struct {
	chrome   *ChromeHandler
	ifShowed bool
}

func NewHTTPChromeHandler() *HTTPChromeSwitch {
	return &HTTPChromeSwitch{NewChromeHandler(), false}
}

func (c *HTTPChromeSwitch) Show(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.RemoteAddr, "127.0.0.1") {
		w.WriteHeader(403)
		fmt.Fprintf(w, "Forbidden")
		return
	}

	if c.ifShowed {
		c.chrome.Cancel()
	}
	c.chrome = NewChromeHandler()
	c.chrome.Show()
	c.ifShowed = true

	SendJSON(w, HTTPResponseFormat{Status: "OK"}, 200)

	return
}

func (c *HTTPChromeSwitch) Off(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.RemoteAddr, "127.0.0.1") {
		w.WriteHeader(403)
		fmt.Fprintf(w, "Forbidden")
		return
	}

	if c.ifShowed {
		c.chrome.Cancel()
	}

	SendJSON(w, HTTPResponseFormat{Status: "OK"}, 200)
	return
}
