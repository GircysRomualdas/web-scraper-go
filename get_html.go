package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "BootCwawler/1.0")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("got network error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("unexpected content type: %s", contentType)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}
	return string(body), nil
}
