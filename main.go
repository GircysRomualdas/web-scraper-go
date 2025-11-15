package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]

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
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
	}

	cfg.crawlPage(cfg.baseURL.String())
	cfg.wg.Wait()
}
