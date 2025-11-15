package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage <url> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("maxConcurrency must be an integer")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("maxPages must be an integer")
		os.Exit(1)
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl: %s\n", baseURL)

	cfg := &config{
		pages:              make(map[string]PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	if err := writeCSVReport(cfg.pages, "report.csv"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
