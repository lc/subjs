# subjs
## Description
A tool to get javascript files from a list of URLS or subdomains. Analyzing javascript files can help you find undocumented endpoints, etc.

It's recommended to pair this with [https://github.com/GerbenJavado/LinkFinder](https://github.com/GerbenJavado/LinkFinder)


## Usage:

`cdl@doggos ~> cat urls.txt | subjs`

[![asciicast](https://asciinema.org/a/225809.svg)](https://asciinema.org/a/225809)


### Save output to a file:

COMING SOON -> `subjs -o output.txt`

## Installation

### Install Command and Download Source With Go Get

`subjs` command will be installed to ```$GOPATH/bin``` and the source code (from `https://github.com/C0RB3N/subjs`) will be found in `$GOPATH/src/github.com/C0RB3N/subjs` with:

```bash
~ ❯ go get -u github.com/C0RB3N/subjs
```

### Install from Github Source

```
~ ❯ git clone https://github.com/C0RB3N/subjs
~ ❯ cd subjs
~ ❯ chmod +x install.sh && ./install.sh
```

## Useful?

<a href="http://buymeacoff.ee/cdl" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
