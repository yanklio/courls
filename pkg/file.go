package scrapper

import (
	"fmt"
	"os"
)

func (s *scraper) setupFileOutput() error {
	if !s.props.isFile {
		return nil
	}

	file, err := os.Create(s.props.FileName)
	if err != nil {
		return err
	}

	s.file = file
	return nil
}

func (s *scraper) writeToFile(count int, statusCode int, url string) {
	if s.file != nil {
		fmt.Fprintf(s.file, "%5d   %3d     %s\n", count, statusCode, url)
	}
}

func (s *scraper) closeFile() {
	if s.file != nil {
		s.file.Close()
	}
}
