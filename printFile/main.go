package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Get file name through command line
	fmt.Println(os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
