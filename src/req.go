package main

import (
	"./link"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run req.go <url>")
		return
	}

	link, err := link.Build(0, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to build URL: %s\n", err)
		return
	}

	fmt.Println(link)
}
