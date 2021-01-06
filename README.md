# subjs
[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![Go ReportCard](https://goreportcard.com/badge/github.com/virenpawar/subjs)](https://goreportcard.com/report/github.com/virenpawar/subjs)

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
$ GO111MODULE=on go get -u -v github.com/virenpawar/subjs
```

### From Binary:
You can download the pre-built [binaries](https://github.com/virenpawar/subjs/releases/) from the releases page and then move them into your $PATH.

```
$ tar xvf subjs_1.0.0_linux_amd64.tar.gz
$ mv subjs /usr/bin/subjs
```

## Changes by me
- Removed unnecessary "subjs" user-agent which was by default.
- Added regex to filter and find more js files. 

## Useful / Suggestions?

- Feel free to ping me on [Twitter](https://twitter.com/VirenPawar_).

---

## Once again thanks to [Corben Leo](https://github.com/lc) for his efforts and his amazing tools.

