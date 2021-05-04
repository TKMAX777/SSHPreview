package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 10 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type HTTPFileNameSender struct {
	conn          *websocket.Conn
	send          chan string
	hasConnection bool
}

// FileNameSender websocket connection on local machine
func NewFileNameSender(message chan string) *HTTPFileNameSender {
	return &HTTPFileNameSender{nil, message, false}
}

func (h *HTTPFileNameSender) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.RemoteAddr, "127.0.0.1") {
		w.WriteHeader(403)
		fmt.Fprintf(w, "Forbidden")
		return
	}

	if h.hasConnection {
		h.conn.WriteJSON(HTTPResponseFormat{Status: "GetNewConnection"})
		h.conn.Close()
		h.hasConnection = false
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		fmt.Println(err)
		return
	}

	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	h.hasConnection = true
	h.conn = conn

	go h.sendMessage()
}

func (h *HTTPFileNameSender) sendMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		h.conn.Close()
	}()

	for {
		select {
		case message, ok := <-h.send:
			h.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				h.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := h.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			fmt.Fprint(w, message)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			h.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := h.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
