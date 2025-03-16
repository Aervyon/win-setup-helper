package main

import (
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/hashicorp/go-getter/v2"
)

func HandleDownloading(cwd string) {
	dir := path.Join(cwd, "win_downloads")

	downloads := strings.Lines(downloadsFile)
	var HttpGetter = &getter.HttpGetter{
		Netrc:                 false,
		XTerraformGetDisabled: true,
		HeadFirstTimeout:      10 * time.Second,
		ReadTimeout:           time.Hour * 24, // For slow internet
	}
	var GitGetter = &getter.GitGetter{
		Detectors: []getter.Detector{
			new(getter.GitHubDetector),
			new(getter.GitDetector),
			new(getter.BitBucketDetector),
			new(getter.GitLabDetector),
		},
	}
	client := &getter.Client{
		DisableSymlinks: true,
		Getters:         []getter.Getter{GitGetter, HttpGetter},
		Decompressors:   getter.DefaultClient.Decompressors,
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
			ProgressListener: &ProgressBar{}, // To implement
		}
		downloadReq(client, req)
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
