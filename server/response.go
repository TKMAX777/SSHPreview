package main

import (
	"encoding/json"
	"net/http"
)

type HTTPResponseFormat struct {
	Status  string
	Content interface{} `json:"Content,omitempty"`
}

// SendJSON write src data to http.ResponseWriter
func SendJSON(w http.ResponseWriter, src interface{}, StatusCode int) (err error) {
	b, err := json.MarshalIndent(src, "", "    ")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	if StatusCode != 0 {
		w.WriteHeader(StatusCode)
	}

	w.Write(b)

	return
}
