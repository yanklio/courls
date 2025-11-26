package scraper

import "os"

type scraper struct {
	props   *scraperProps
	results chan<- *CompletedUrl
	count   int
	file    *os.File
}

type scraperProps struct {
	Url string
	Limit int

	isFile bool
	FileName string
}

func NewScraperProps(url string, limit int, fileName string) *scraperProps {
	return &scraperProps{
		Url:    url,
		Limit:  limit,
		FileName:   fileName,
		isFile: fileName != "",
	}
}

type CompletedUrl struct {
	Id int
	StatusCode int
	Url string
}

func NewCompletedUrl(id int, statusCode int, url string) *CompletedUrl {
	return &CompletedUrl{
		Id:         id,
		StatusCode: statusCode,
		Url:        url,
	}
}

type scraperResult struct {
	count int
}

func NewScraperResult(count int) *scraperResult {
	return &scraperResult{
		count: count,
	}
}
