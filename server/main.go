package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

const SettingFileName = "settings.json"

func init() {
	Settings.ListenPort = os.Getenv("PreviewListenPort")
	if Settings.ListenPort == "" {
		panic("Listen port is not specified")
	}

	b, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic("panic: index file not found")
	}

	b = bytes.Replace(b, []byte("{{$ADDRESS}}"), []byte("127.0.0.1:"+Settings.ListenPort), 1)
	err = ioutil.WriteFile(filepath.Join("resources", "index.html"), b, 0666)
	if err != nil {
		panic("panic: could not make index file")
	}

}

func main() {
	NewHTTPHandler().Start()
}
