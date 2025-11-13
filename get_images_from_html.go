package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var urls []string

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}
		src = strings.TrimSpace(src)
		if src == "" {
			return
		}

		rel, err := url.Parse(src)
		if err != nil {
			return
		}

		url := baseURL.ResolveReference(rel).String()
		urls = append(urls, url)
	})

	return urls, nil
}
