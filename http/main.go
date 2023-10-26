package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// It will initialize the byte slice with the particular size 99999
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}
