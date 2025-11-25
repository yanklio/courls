package scrapper

import "os"

type scrapperProps struct {
	Url string
	Limit int

	isFile bool
	File *os.File
}

func NewScrapperProps(url string, limit int, fileName string) *scrapperProps {

	var file *os.File
	if fileName != "" {
		var err error
		file, err = os.Create(fileName)
		if err != nil {
			panic(err)
		}
	}

	return &scrapperProps{
		Url:    url,
		Limit:  limit,
		File:   file,
		isFile: fileName != "",
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
