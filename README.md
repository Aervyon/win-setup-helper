# win-setup-helper
A tool to help setup a Windows install by downloading things and going to webpages.

Currently, you need to build it yourself.

# Building From Source

You need [go](https://go.dev/dl) 1.24 or higher
Next, clone this repository
```pwsh
git clone https://github.com/Aervyon/win-setup-helper
```

Make a `download.txt` and fill it with a bunch of URLs for webpages you'd like to download, you can also supply Git repositories.

Each webpage or repository needs to be on a new line, and repositories cannot start with https or http.
Example:
```txt
https://github.com/git-for-windows/git/releases/download/v2.48.1.windows.1/Git-2.48.1-64-bit.exe
github.com/Aervyon/win-setup-helper
https://go.dev/dl/go1.24.1.windows-amd64.msi
```

Test to ensure it works
Build with
```pwsh
go build . -o out/setup-downloader.exe
```
and place it somewhere you can easily access it (another drive, a usb stick, SD card, etc)

To run this, simply navigate to where you stored it and do
```pwsh
./setup-downloader.exe
```
in PowerShell or Command

## Other Plans
- Open pages in the browser
- Automatically run the installers it downloads
- Install directly from winget
- Post-install configuration with a JSON or TOML file.

# License
It's MIT licensed, for now.
Go nuts.