package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	parsedURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if currentURL.Hostname() != parsedURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, ok := pages[normalizedURL]; ok {
		pages[normalizedURL] += 1
		return
	}

	pages[normalizedURL] = 1
	fmt.Printf("crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	nextURLs, err := getURLsFromHTML(html, parsedURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, nextURL := range nextURLs {
		crawlPage(rawBaseURL, nextURL, pages)
	}
}
