package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// go routine
		go checkLink(link, c)
	}

	// infinite loop
	// for {
	// 	// for i := 0; i < len(links); i++ {
	// 	// fmt.Println(<-c)
	// 	go checkLink(<-c, c)
	// }

	// alternative syntax to above (initite loop)
	// note: range c = watch channel c
	for l := range c {
		// function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//send information to channel
		// c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")

	//send information to channel
	// c <- "Yep its up"
	c <- link
}
