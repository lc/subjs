package subjs

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const version = `1.0.2`

type SubJS struct {
	client *http.Client
	opts   *Options
}

func New(opts *Options) *SubJS {
	c := &http.Client{
		Timeout:   time.Duration(opts.Timeout) * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	return &SubJS{client: c, opts: opts}
}
func (s *SubJS) Run() error {
	// setup input
	var input *os.File
	var err error
	// if input file not specified then read from stdin
	if s.opts.InputFile == "" {
		input = os.Stdin
	} else {
		// otherwise read from file
		input, err = os.Open(s.opts.InputFile)
		if err != nil {
			return fmt.Errorf("Could not open input file: %s", err)
		}
		defer input.Close()
	}

	// init channels
	urls := make(chan string)
	results := make(chan string)

	// start workers
	var w sync.WaitGroup
	for i := 0; i < s.opts.Workers; i++ {
		w.Add(1)
		go func() {
			s.fetch(urls, results)
			w.Done()
		}()
	}
	// setup output
	var out sync.WaitGroup
	out.Add(1)
	go func() {
		for result := range results {
			fmt.Println(result)
		}
		out.Done()
	}()
	scan := bufio.NewScanner(input)
	for scan.Scan() {
		u := scan.Text()
		if u != "" {
			urls <- u
		}
	}
	close(urls)
	w.Wait()
	close(results)
	out.Wait()
	return nil
}
func (s *SubJS) fetch(urls <-chan string, results chan string) {
	for u := range urls {
		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			continue
		}
		if s.opts.UserAgent != "" {
			req.Header.Add("User-Agent", s.opts.UserAgent)
		}
		resp, err := s.client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			continue
		}
		u, err := url.Parse(u)
		if err != nil {
			// handle error
			//log.Fatalf("error parsing url: %v", err)
			return
		}
		doc.Find("script").Each(func(index int, s *goquery.Selection) {
			js, _ := s.Attr("src")
			if js != "" {
				if strings.HasPrefix(js, "http://") || strings.HasPrefix(js, "https://") {
					results <- js
				} else if strings.HasPrefix(js, "//") {
					js := fmt.Sprintf("%s:%s", u.Scheme, js)
					results <- js
				} else if strings.HasPrefix(js, "/") {
					js := fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, js)
					results <- js
				} else {
					js := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, js)
					results <- js
				}
			}
			r := regexp.MustCompile(`[(\w./:)]*js`)
			matches := r.FindAllString(s.Contents().Text(), -1)
			for _, js := range matches {
				if strings.HasPrefix(js, "//") {
					js := fmt.Sprintf("%s:%s", u.Scheme, js)
					results <- js
				} else if strings.HasPrefix(js, "/") {
					js := fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, js)
					results <- js
				}
			}
		})
		doc.Find("div").Each(func(index int, s *goquery.Selection) {
			js, _ := s.Attr("data-script-src")
			if js != "" {
				if strings.HasPrefix(js, "http://") || strings.HasPrefix(js, "https://") {
					results <- js
				} else if strings.HasPrefix(js, "//") {
					js := fmt.Sprintf("%s:%s", u.Scheme, js)
					results <- js
				} else if strings.HasPrefix(js, "/") {
					js := fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, js)
					results <- js
				} else {
					js := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, js)
					results <- js
				}
			}
		})
	}
}
