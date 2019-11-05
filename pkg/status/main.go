package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.nl",
		"http://reddit.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.de",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		//  Make use of pass by value feature in golang
		//  Dont access a variable from the child routine
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		// go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "something doesn't seem right!")
		c <- link
		return
	}

	fmt.Println(link, "is all good!")
	c <- link
}
