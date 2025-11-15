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
starting crawl: https://blog.boot.dev/
crawling https://blog.boot.dev/
crawling https://blog.boot.dev/privacy/
crawling https://blog.boot.dev/news/bootdev-beat-2025-11/
crawling https://blog.boot.dev/tutorials/python/loops/
crawling https://blog.boot.dev/education/vibe-coding-hell/
crawling https://blog.boot.dev/news/bootdev-beat-2025-10/
crawling https://blog.boot.dev/news/bootdev-beat-2025-09/
crawling https://blog.boot.dev/create-a-course/
crawling https://blog.boot.dev/news/training-grounds-launch/
crawling https://blog.boot.dev/news/bootdev-beat-2025-08/
crawling https://blog.boot.dev/news/hackathon-2025/
crawling https://blog.boot.dev/news/bootdev-beat-2025-07/
crawling https://blog.boot.dev/news/bootdev-beat-2025-06/
crawling https://blog.boot.dev/education/is-boot-dev-free/
crawling https://blog.boot.dev/news/bootdev-beat-2025-05/
crawling https://blog.boot.dev/news/bootdev-beat-2025-04/
crawling https://blog.boot.dev/news/bootdev-beat-2025-03/
crawling https://blog.boot.dev/news/bootdev-beat-2025-02/
crawling https://blog.boot.dev/computer-science/18-months-with-gpt-4/
crawling https://blog.boot.dev/news/bootdev-beat-2025-01/
crawling https://blog.boot.dev/tutorials/python/lists/
crawling https://blog.boot.dev/tutorials/python/functions/
crawling https://blog.boot.dev/tutorials/python/variables/
crawling https://blog.boot.dev/news/bootdev-beat-2024-12/
crawling https://blog.boot.dev/news/bootdev-beat-2024-11/
```
