package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	file        = flag.String("i", "", "input file containg urls")
	format      = flag.Bool("json", false, "output in json format")
	wayback     = flag.Bool("wayback", false, "retrieve javascript files from the wayback machine")
	urls        []string
	input       io.Reader
	seenWayback map[string]bool
	waybackresp [][]string
)

var subjs = &http.Client{
	Timeout: time.Second * 20,
}

func main() {
	flag.Parse()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	output := make(map[string][]string)
	if *file == "" {
		input = os.Stdin
	} else {
		infp, err := os.Open(*file)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		input = infp
	}
	s := bufio.NewScanner(input)
	for s.Scan() {
		urls = append(urls, s.Text())
	}
	if s.Err() != nil {
		log.Fatalf("Error retrieving input: %v", s.Err())
	}
	for _, host := range urls {
		found := getScripts(host)
		if *wayback {
			u, err := url.Parse(host)
			if err != nil {
				log.Fatalf("Error parsing url: %v", err)
			}
			temp := waybackUrls(u.Hostname())
			for _, js := range temp {
				found = append(found, js)
			}
		}
		for _, js := range dedupe(found) {
			output[host] = append(output[host], js)
		}
	}
	if *format {
		bytes, err := json.MarshalIndent(output, "", "    ")
		if err != nil {
			log.Fatalf("Error JSON Marshalling data: %v", err)
		}
		fmt.Println(string(bytes))
	} else {
		for _, items := range output {
			for _, file := range items {
				fmt.Println(file)
			}
		}
	}
}
func getScripts(domain string) []string {
	var found []string
	resp, err := subjs.Get(domain)
	if err != nil {
		return found
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error parsing response from: ", domain)
	}
	u, err := url.Parse(domain)
	if err != nil {
		log.Fatalf("error parsing url: %v", err)
	}
	doc.Find("script").Each(func(index int, s *goquery.Selection) {
		js, _ := s.Attr("src")
		if js != "" {
			if strings.HasPrefix(js, "http://") || strings.HasPrefix(js, "https://") {
				found = append(found, js)
			} else if strings.HasPrefix(js, "//") {
				js := fmt.Sprintf("%s:%s", u.Scheme, js)
				found = append(found, js)
			} else {
				js := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, js)
				found = append(found, js)
			}
		}
	})
	doc.Find("div").Each(func(index int, s *goquery.Selection) {
		js, _ := s.Attr("data-script-src")
		if js != "" {
			if strings.HasPrefix(js, "http://") || strings.HasPrefix(js, "https://") {
				found = append(found, js)
			} else if strings.HasPrefix(js, "//") {
				js := fmt.Sprintf("%s:%s", u.Scheme, js)
				found = append(found, js)
			} else if strings.HasPrefix(js, "/") {
				js := fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, js)
				found = append(found, js)
			} else {
				js := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, js)
				found = append(found, js)
			}
		}

	})
	return found
}
func waybackUrls(hostname string) []string {
	var found []string
	tg := fmt.Sprintf("http://web.archive.org/cdx/search/cdx?url=%s/*&output=json&collapse=urlkey&fl=original", hostname)
	r, err := subjs.Get(tg)
	if err != nil {
		log.Printf("Error in http request: %v\n", err)
		return found
	}
	defer r.Body.Close()
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v\n", err)
		return found
	}
	err = json.Unmarshal(resp, &waybackresp)
	if err != nil {
		log.Printf("Error unmarshalling response: %v\n", err)
	}
	first := true
	for _, result := range waybackresp {
		if first {
			// skip first result from wayback machine
			// always is "original"
			first = false
			continue
		}
		u, err := url.Parse(result[0])
		if err != nil {
			continue
		}
		if strings.HasSuffix(u.Path, ".js") {
			found = append(found, result[0])
		}
	}
	return found
}
func dedupe(all []string) []string {
	seen := make(map[string]bool)
	unique := []string{}
	for b := range all {
		if !seen[all[b]] {
			seen[all[b]] = true
			unique = append(unique, all[b])
		}
	}
	return unique
}
