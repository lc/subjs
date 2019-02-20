package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/C0RB3N/subjs/banner"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	var (
		outf         string
		domains      []string
		singleDomain string
	)

	singleDomainOut := make(map[string][]string)
	out := make(map[string][]string)

	//menu
	fmt.Println(banner.Banner())
	flag.StringVar(&outf, "o", "", "Name of the output file")
	flag.StringVar(&singleDomain, "d", "", "Name of the uniq domain to search for")
	flag.Parse()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	var subjs = &http.Client{
		Timeout: time.Second * 5,
	}

	re := regexp.MustCompile("^(?:https?:\\/\\/)?(?:www\\.)?([^\\/]+)")

	// scan single domain
	if outf != "" && singleDomain != "" {

		resp, err := subjs.Get(singleDomain)
		host := re.FindStringSubmatch(singleDomain)

		if err == nil {
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				fmt.Println("Error parsing response from: ", singleDomain)
			}

			doc.Find("script").Each(func(index int, s *goquery.Selection) {
				js, _ := s.Attr("src")
				if js != "" {
					if strings.HasPrefix(js, "http://") || strings.HasPrefix(js, "https://") || strings.HasPrefix(js, "//") {
						singleDomainOut[singleDomain] = append(singleDomainOut[singleDomain], js)
					} else {
						js := strings.Join([]string{host[1], js}, "")
						singleDomainOut[singleDomain] = append(singleDomainOut[singleDomain], js)
					}
				}
			})

			if len(singleDomainOut) != 0 {
				bytes, err := json.MarshalIndent(singleDomainOut, "", "    ")
				if err == nil {
					fmt.Println(string(bytes))
				}
				if outf != "" {
					ioutil.WriteFile(outf, bytes, 0644)
				}
			}
			fmt.Println("[+] Operation sucess ouput in: ", outf)
		}
	} else {

		// validation open file
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

		// send urls from file to http handler
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
			// creation file output
			if len(out) != 0 {
				bytes, err := json.MarshalIndent(out, "", "    ")
				if err == nil {
					fmt.Println(string(bytes))
				}
				if outf != "" {
					ioutil.WriteFile(outf, bytes, 0644)
				}
			}
			fmt.Println("[+] Operation sucess ouput in: ", outf)
		}
	}

} //end main
