package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	var subjs = &http.Client{
		Timeout: time.Second * 5,
	}

	re := regexp.MustCompile("^(?:https?:\\/\\/)?(?:www\\.)?([^\\/]+)")
	var domains []string
	out := make(map[string][]string)
	m, _ := os.Stdin.Stat()
	if (m.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			domains = append(domains, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "-> subjs - corben leo\n-> usage: cat urls.txt | subjs", err)
			os.Exit(3)
		}
	} else {
		fmt.Fprintf(os.Stderr, "-> subjs - corben leo \n-> usage: cat urls.txt | subjs")
	}
	for _, domain := range domains {
		resp, err := subjs.Get(domain)

		host := re.FindStringSubmatch(domain)
		if err == nil {
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				fmt.Println("Error parsing response from: ", domain)
			}
			doc.Find("script").Each(func(index int, s *goquery.Selection) {
				js, _ := s.Attr("src")
				if js != "" {
					if strings.HasPrefix(js, "http://") || strings.HasPrefix(js, "https://") || strings.HasPrefix(js, "//") {
						out[domain] = append(out[domain], js)
					} else {
						js := strings.Join([]string{host[1], js}, "")
						out[domain] = append(out[domain], js)
					}
				}
			})
		}
	}
	if len(out) != 0 {
		bytes, err := json.MarshalIndent(out, "", "    ")
		if err == nil {
			fmt.Println(string(bytes))
		}
	}
}
