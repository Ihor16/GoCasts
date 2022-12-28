package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	lw := logWriter{}
	_, err = io.Copy(lw, resp.Body)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Printf("%s", bs)
	fmt.Printf("Wrote %d bytes\n", len(bs))
	return len(bs), nil
}
