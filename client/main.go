package client

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func init() {
	Settings.ListenSock = os.Getenv("PreviewListenSock")
	if Settings.ListenSock == "" {
		Settings.ListenPort = os.Getenv("PreviewListenPort")
		if Settings.ListenPort == "" {
			panic("both of ListenPort and LitenSock not specified")
		}
	}
}

func Start() {

	flag.Parse()
	var args = flag.Args()

	var req *RequestHandler
	if Settings.ListenSock == "" {
		req = NewRequestHandler("localhost:"+Settings.ListenPort, SockTypeTCP)
	} else {
		req = NewRequestHandler(Settings.ListenSock, SockTypeUNIX)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var enter = make(chan bool, 1)
	var back = make(chan bool)

	go func() {
		for {
			var scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()

			switch strings.TrimSpace(scanner.Text()) {
			case "back", "b":
				back <- true
			default:
				enter <- true
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	err := req.ChromeShow()
	if err != nil {
		fmt.Printf("Chrome excute error:\n%s\n", err.Error())
		return
	}

	for i := 0; i < len(args); i++ {
		var p = args[i]

		fmt.Println(p)

		err = req.Verify(filepath.Base(p))
		if err != nil {
			fmt.Println(err.Error())
			goto wait
		}

		{
			f, err := os.Open(p)
			if err != nil {
				continue
			}

			req.Send(HTTPPreviewData{
				Data:     f,
				FileName: f.Name(),
			})

			f.Close()
		}
	wait:
		for {
			select {
			case <-interrupt:
				req.ChromeOff()
				goto end
			case <-enter:
				break wait
			case <-back:
				if i > 1 {
					i -= 2
				}
				break wait
			case <-ticker.C:
			}
		}
	}
end:
	req.ChromeOff()
	fmt.Println()

}
