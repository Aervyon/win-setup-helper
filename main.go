package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/hashicorp/go-getter/v2"
)

//go:embed download.txt
var downloadsFile string

func main() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dir = path.Join(dir, "win_downloads")

	downloads := strings.Lines(downloadsFile)
	var httpGetter = &getter.HttpGetter{
		ReadTimeout: time.Hour * 24, // For slow internet
	}
	var gitGetter = &getter.GitGetter{
		Timeout: time.Hour * 1, // For massive projects, like Chromium
	}
	client := &getter.Client{
		DisableSymlinks: true,
		Getters:         []getter.Getter{httpGetter, gitGetter},
	}

	for download := range downloads {
		download, _ = strings.CutSuffix(download, "\n")
		dest := strings.SplitAfter(download, "?")[0]
		dest = path.Base(dest)
		dest = strings.TrimSuffix(dest, path.Ext(dest))
		fmt.Println(dest)
		req := &getter.Request{
			Src:              download,
			Dst:              path.Join(dir, dest),
			GetMode:          getter.ModeAny,
			ProgressListener: &progressBar{}, // To implement
		}
		downloadReq(client, req)
	}

	// Define the source and destination paths
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func downloadReq(client *getter.Client, req *getter.Request) {
	src := strings.Split(req.Src, "/")
	fmt.Printf("Downloading %v\n", req.Src)
	result, err := client.Get(context.Background(), req)
	if err != nil {
		fmt.Printf("Downloading error for %v occurred: %v\n", src[len(src)-1], err.Error())
		return
	}
	fmt.Printf("Downloaded %v to %v\n", src[len(src)-1], result.Dst)
}
