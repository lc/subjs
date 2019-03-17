package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var oJson bool
	flag.BoolVar(&oJson, "json", false, "Format output as json")
	flag.Parse()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	var subjs = &http.Client{
		Timeout: time.Second * 20,
	}

	var domains []string
	out := make(map[string][]string)
	var files []string
	m, _ := os.Stdin.Stat()
	if (m.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			domains = append(domains, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "-> subjs - corben leo\n-> usage: cat urls.txt | subjs\n\nError: %v", err)
			os.Exit(3)
		}
	} else {
		fmt.Fprintf(os.Stderr, "-> subjs - corben leo \n-> usage: cat urls.txt | subjs")
	}
	for _, domain := range domains {
		resp, err := subjs.Get(domain)

		if err != nil {
			log.Fatalf("Error making HTTP request: %v", err)
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("Error parsing response from: ", domain)
		}
		doc.Find("script").Each(func(index int, s *goquery.Selection) {
			js, _ := s.Attr("src")
			if js != "" {
				if strings.HasPrefix(js, "http://") || strings.HasPrefix(js, "https://") || strings.HasPrefix(js, "//") {
					Output(out, &files, domain, js, oJson)
				} else {
					u, err := url.Parse(domain)
					if err != nil {
						log.Fatalf("Error parsing domain: %v", err)
					}
					js := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, js)
					Output(out, &files, domain, js, oJson)
				}
			}
		})

	}
	if len(out) != 0 || len(files) != 0 {
		if oJson {
			bytes, err := json.MarshalIndent(out, "", "    ")
			if err != nil {
				log.Fatalf("Error JSON Marshalling data: %v", err)
			}
			fmt.Println(string(bytes))
		} else {
			for _, file := range files {
				fmt.Println(file)
			}
		}
	}
}
func Output(out map[string][]string, files *[]string, domain string, js string, isJson bool) {
	if isJson {
		out[domain] = append(out[domain], js)
	} else {
		*files = append(*files, js)
	}
}
