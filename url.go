package main

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

// Stolen from https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8?permalink_comment_id=5084817#gistcomment-5084817
func openURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "powershell"
		args = []string{"-Command", "Start-Process", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		// Check if running under WSL
		if isWSL() {
			// Use 'cmd.exe /c start' to open the URL in the default Windows browser
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			// Use xdg-open on native Linux environments
			cmd = "xdg-open"
			args = []string{url}
		}
	}
	if len(args) > 1 {
		// args[0] is used for 'start' command argument, to prevent issues with URLs starting with a quote
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}
	return exec.Command(cmd, args...).Start()
}

// isWSL checks if the Go program is running inside Windows Subsystem for Linux
func isWSL() bool {
	releaseData, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}

func HandleOpening() {
	urls := strings.Lines(openFile)
	for urlEntry := range urls {
		parsed, err := url.Parse(urlEntry)
		if err != nil {
			fmt.Printf("Error parsing URL %v: %v\n", urlEntry, err)
			continue
		}
		if parsed.Scheme == "" {
			fmt.Printf("WARNING: URL %v does not have a scheme. Defaulting to https://\nPlease add your URLs with a scheme next time\n", urlEntry)
			urlEntry = "https://" + urlEntry
		}
		err = openURL(urlEntry)
		if err != nil {
			fmt.Println(err)
		}
	}
}
