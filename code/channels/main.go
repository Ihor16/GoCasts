package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
	}
	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	for link := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkStatus(l, c)
		}(link)
	}
}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
