package scraper

import "os"

// scraper holds the state of a scraping process.
type scraper struct {
	props   *scraperProps
	results chan<- *CompletedUrl
	count   int
	file    *os.File
}

// scraperProps contains the configuration for a scraping process.
type scraperProps struct {
	Url      string
	Limit    int
	isFile   bool
	FileName string
}

// NewScraperProps creates a new scraperProps configuration.
func NewScraperProps(url string, limit int, fileName string) *scraperProps {
	return &scraperProps{
		Url:      url,
		Limit:    limit,
		FileName: fileName,
		isFile:   fileName != "",
	}
}

// CompletedUrl represents a URL that has been scraped.
type CompletedUrl struct {
	Id         int
	StatusCode int
	Url        string
}

// NewCompletedUrl creates a new CompletedUrl instance.
func NewCompletedUrl(id int, statusCode int, url string) *CompletedUrl {
	return &CompletedUrl{
		Id:         id,
		StatusCode: statusCode,
		Url:        url,
	}
}

// scraperResult holds the final result of a scraping process.
type scraperResult struct {
	count int
}

// NewScraperResult creates a new scraperResult instance.
func NewScraperResult(count int) *scraperResult {
	return &scraperResult{
		count: count,
	}
}
