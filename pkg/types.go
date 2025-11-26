package scrapper

import "os"

type scraper struct {
	props   *scrapperProps
	results chan<- *CompletedUrl
	count   int
	file    *os.File
}

type scrapperProps struct {
	Url string
	Limit int

	isFile bool
	FileName string
}

func NewScrapperProps(url string, limit int, fileName string) *scrapperProps {
	return &scrapperProps{
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

type scrapperResult struct {
	count int
}

func NewScrapperResult(count int) *scrapperResult {
	return &scrapperResult{
		count: count,
	}
}
