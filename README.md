# Web scraper

A concurrent web crawler built in Go that scrapes pages and generates a CSV report.

This is the starter code used in Boot.dev's [Build a Web Scraper in Go](https://www.boot.dev/courses/build-web-scraper-golang) course.

---

## Requirements

- Go 1.20+

---

## Installation

1. Clone the repository.

2. Install dependencies:
```bash
go mod download
```

## Usage

Run web scraper:
```bash
go run . <url> <maxConcurrency> <maxPages>
```

- `<url>` The website address where the crawler begins its search.
- `<maxConcurrency>` Number of pages the crawler can fetch at the same time.
- `<maxPages>` Total number of pages to scrape before stopping.

### Example
```bash
go run . "https://blog.boot.dev/" 3 25
```

Output:
```bash
Starting crawl on: https://blog.boot.dev/
Crawling page: https://blog.boot.dev/
Crawling page: https://blog.boot.dev/privacy/
Crawling page: https://blog.boot.dev/create-a-course/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2024-12/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2025-04/
Crawling page: https://blog.boot.dev/news/training-grounds-launch/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2025-08/
Crawling page: https://blog.boot.dev/news/hackathon-2025/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2025-07/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2025-06/
Crawling page: https://blog.boot.dev/education/is-boot-dev-free/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2025-05/
Crawling page: https://blog.boot.dev/page/2/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2024-11/
Crawling page: https://blog.boot.dev/education/state-of-learning-to-code-2024/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2024-10/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2025-01/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2025-03/
Crawling page: https://blog.boot.dev/news/bootdev-beat-2025-02/
Crawling page: https://blog.boot.dev/computer-science/18-months-with-gpt-4/
Crawling page: https://blog.boot.dev/contact/
Crawling page: https://blog.boot.dev/categories/
Crawling page: https://blog.boot.dev/tutorials/python/loops/
Crawling page: https://blog.boot.dev/tutorials/python/lists/
Crawling page: https://blog.boot.dev/education/vibe-coding-hell/
CSV report generated successfully: report.csv
```
