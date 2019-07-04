package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func craw(path string, url string) {
	req, _ := http.NewRequest("GET", url+path, nil)
	client := &http.Client{}
	response, _ := client.Do(req)
	tokenizer := html.NewTokenizer(response.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()
		if tokenType == html.StartTagToken { //opening tag
			switch token.DataAtom.String() {
			case "a", "link": //link tags
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						fmt.Println(attr.Val)

					}
				}
			}
		}
	}
}

func main() {
	link := "http://monzo.com"
	// link := "http://locahost:8000"
	// req, _ := http.NewRequest("GET", "http://localhost:8000", nil)
	req, _ := http.NewRequest("GET", link, nil)
	client := &http.Client{}
	response, _ := client.Do(req)
	tokenizer := html.NewTokenizer(response.Body)
	//body, err := ioutil.ReadAll(io.LimitReader(response.Body, 1048576))
	//tokenizer.
	// for {
	// 	token := tokenizer.Next()
	// 	fmt.Println(token.String(), tokenizer.Token().Data, tokenizer.Token().DataAtom, tokenizer.Token().Attr)
	// 	time.Sleep(1 * time.Second)
	// }
	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					//fmt.Println(tl, attr.Val)
					urlFinal := fixURL(tl, link)
					if strings.Contains(urlFinal, link) {
						fmt.Println(urlFinal, tl)
					}
				}
			}
		}
	}
}

func fixURL(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseURL.ResolveReference(uri)
	return uri.String()
}

// trimHash slices a hash # from the link
func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}
