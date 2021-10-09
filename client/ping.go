package client

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func KeepPing() {
	if Settings.ListenSock == "" {
		panic("Error: ListenSock not specified")
	}

	checkInterval, err := strconv.Atoi(os.Getenv("PreviewCheckInterval"))
	if err != nil {
		checkInterval = 5
	}
	var req = NewRequestHandler(Settings.ListenSock, SockTypeUNIX)

	fmt.Printf("Start checking: %s", filepath.Join(Settings.ListenSock, "http.sock"))

	for {
		time.Sleep(time.Second * time.Duration(checkInterval))
		_, err := os.Stat(filepath.Join(Settings.ListenSock, "http.sock"))
		if err != nil {
			continue
		}

		err = req.Ping()
		if err != nil {
			os.Remove(filepath.Join(Settings.ListenSock, "http.sock"))
		}
	}
}
