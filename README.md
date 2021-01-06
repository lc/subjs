# subjs
[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![Go ReportCard](https://goreportcard.com/badge/github.com/lc/gau)](https://goreportcard.com/report/github.com/lc/subjs)

subjs fetches javascript files from a list of URLS or subdomains. Analyzing javascript files can help you find undocumented endpoints, secrets, and more.

It's recommended to pair this with [gau](https://github.com/lc/gau) and then [https://github.com/GerbenJavado/LinkFinder](https://github.com/GerbenJavado/LinkFinder)

# Resources
- [Usage](#usage)
- [Installation](#installation)

## Usage:
Examples:
```bash
$ cat urls.txt | subjs 
$ subjs -i urls.txt
$ cat hosts.txt | gau | subjs
```

To display the help for the tool use the `-h` flag:

```bash
$ subjs -h
```

| Flag | Description | Example |
|------|-------------|---------|
| `-c` | Number of concurrent workers | `subjs -c 40` |
| `-i` | Input file containing URLS | `subjs -i urls.txt` |
| `-t` | Timeout (in seconds) for http client (default 15) | `subjs -t 20` |
| `-ua` | User-Agent to send in requests | `subjs -ua "Chrome..."` |
| `-version` | Show version number | `subjs -version"` |


## Installation
### From Source:

```
$ GO111MODULE=on go get -u -v github.com/lc/subjs@latest
```

### From Binary
You can download the pre-built [binaries](https://github.com/lc/subjs/releases/) from the releases page and then move them into your $PATH.

```
$ tar xvf subjs_1.0.0_linux_amd64.tar.gz
$ mv subjs /usr/bin/subjs
```

## Useful?

<a href="http://buymeacoff.ee/cdl" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
