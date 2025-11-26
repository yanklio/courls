package scraper

import (
	"fmt"
	"os"
)

// setupFileOutput creates and opens the output file if a filename is provided.
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

// writeToFile writes the scraped URL information to the output file.
func (s *scraper) writeToFile(count int, statusCode int, url string) {
	if s.file != nil {
		fmt.Fprintf(s.file, "%5d   %3d     %s\n", count, statusCode, url)
	}
}

// closeFile closes the output file if it is open.
func (s *scraper) closeFile() {
	if s.file != nil {
		s.file.Close()
	}
}
