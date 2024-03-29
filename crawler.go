package main

import (
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func formatURL(base string, link string) string {
	linkURL, err := url.Parse(link)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uriFormatted := baseURL.ResolveReference(linkURL)
	return uriFormatted.String()
}

func returnLocalLinks(baseURL string, links []string) (localLinks []string) {
	var ret []string
	for _, link := range links {
		if strings.HasPrefix(link, baseURL) {
			ret = append(ret, link)
		}
	}
	return ret
}

func isLinkFound(link string, linksFound []string) bool {
	for _, linkFound := range linksFound {
		if link == linkFound {
			return true
		}
	}
	return false
}

func addLinkToList(link string, linksFound []string) []string {
	if !isLinkFound(link, linksFound) {
		linksFound = append(linksFound, link)
	}
	return linksFound
}

func returnAllLinks(url string, linksFound []string) []string {
	var links []string
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	response, _ := client.Do(req)
	tokenizer := html.NewTokenizer(response.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					link := formatURL(url, attr.Val)
					linksFound = addLinkToList(link, linksFound)
				}
			}
		}
	}
}
