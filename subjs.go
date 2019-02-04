package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println(" -- subjs -- \n- by corben leo -")
	re := regexp.MustCompile("^(?:https?:\\/\\/)?(?:www\\.)?([^\\/]+)")
	var domains []string
	m, _ := os.Stdin.Stat()
	if (m.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			domains = append(domains, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "\n-> Usage: cat urls.txt | subjs", err)
			os.Exit(3)
		}
	} else {
		fmt.Fprintf(os.Stderr, "\n-> Usage: cat urls.txt | subjs")
	}
	for _, domain := range domains {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		client := &http.Client{Timeout: 3 * time.Second}
		resp, err := client.Get(domain)
		host := re.FindStringSubmatch(domain)
		if err == nil {
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				fmt.Println("Error parsing response from: ", domain)
			}
			fmt.Println("====", domain, "====")
			doc.Find("script").Each(func(index int, s *goquery.Selection) {
				js, _ := s.Attr("src")
				if js != "" {
					if strings.HasPrefix(js, "http://") || strings.HasPrefix(js, "https://") || strings.HasPrefix(js, "//") {
						fmt.Println(js)
					} else {
						js := strings.Join([]string{host[1], js}, "")
						fmt.Println(js)
					}
				}
			})
		}
	}
}
