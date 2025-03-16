package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed download.txt
var downloadsFile string

//go:embed open.txt
var openFile string

func main() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if downloadsFile != "" {
		HandleDownloading(dir)
	} else {
		fmt.Println("No files need to be downloaded. Skipping.")
	}

	if openFile != "" {
		HandleOpening()
	} else {
		fmt.Println("No URLs need to be opened. Skipping.")
	}
}
