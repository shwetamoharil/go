package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"https://www.google.com/", "https://www.facebook.com/", "https://go.dev/", "https://stackoverflow.com/"}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // we only use go keyword in front of function calls
	}

	// for {
	// 	go checkLink(<-c, c) // continue executing loop
	// }

	// Alternative for loop for channel
	for l := range c { // This means wait for the channel to return some value and once value is returned assign it to l
		go func(l string) { // This is like anonymous func in  JS
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
		// go checkLink(l, c)
	}

	// for i := 0; i < len(links); i++ {

	// 	fmt.Println(<-c) // It will block execution of for loop as it will wait for data into channel
	// }

	// fmt.Println(<-c)
	// This is a blocking call, it will make main routine sleep
	// The go routine that first resolves resp will sent back data in channel
	// As soon as channel gets data, main routine reads it and then exits, hence we see only one link is up
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // Blocking call - until this completes we cannot proceed further
	if err != nil {
		fmt.Println(link, " might be down!")
		//c <- "Might be down"
		c <- link
		return
	}
	//c <- "Link is up"
	c <- link
	fmt.Println(link, " is up!")
}

// sending data into channels

// channel <- 5 (sending data into channel)
// myNumber <- channel (wait for value to be sent into this channel and when you get one assign it to myNumber)
// fmt.Println(<- channel) (wait for value to be sent into the channel and when you get one log it out immediately)
