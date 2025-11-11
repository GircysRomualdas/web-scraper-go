package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	return doc.Find("h1").First().Text()
}

func getFirstParagraphFromHTML(html string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	paragraph := doc.Find("main").Find("p").First().Text()

	if paragraph == "" {
		return doc.Find("p").First().Text()
	}

	return paragraph
}
