package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"}

	if err := writer.Write(header); err != nil {
		return err
	}

	for url, pageData := range pages {
		row := []string{
			url,
			pageData.H1,
			pageData.FirstParagraph,
			strings.Join(pageData.OutgoingLinks, ";"),
			strings.Join(pageData.ImageURLs, ";"),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
