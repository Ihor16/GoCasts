package main

import (
	"io"
	"log"
	"os"
)

func main() {
	name := os.Args[1]

	file, err := os.Open(name)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
