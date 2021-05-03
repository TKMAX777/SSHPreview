package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func init() {
	Settings.ListenPort = os.Getenv("PreviewListenPort")
}

func main() {

	flag.Parse()
	var args = flag.Args()

	var req = NewRequestHandler("http://localhost:" + Settings.ListenPort)

	interrupt := make(chan os.Signal, 1)

	enter := make(chan bool, 1)

	go func() {
		for {
			bufio.NewScanner(os.Stdin).Scan()
			enter <- true
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	req.ChromeShow()

	for _, p := range args {
		f, err := os.Open(p)
		if err != nil {
			continue
		}
		fmt.Println(p)
		req.Send(HTTPPreviewData{
			Data:     f,
			FileName: f.Name(),
		})

		f.Close()

	wait:
		for {
			select {
			case <-interrupt:
				req.ChromeOff()
				os.Exit(1)
			case <-enter:
				break wait
			case <-ticker.C:
			}
		}
	}
	req.ChromeOff()

}
