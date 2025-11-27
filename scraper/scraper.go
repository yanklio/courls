package scraper

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly/v2"
)

// Scrap initiates the web scraping process with the given properties.
// It returns a channel that streams the scraping results.
func Scrap(props *scraperProps) <-chan *CompletedUrl {
	results := make(chan *CompletedUrl)

	go func() {
		defer close(results)

		scraper := &scraper{
			props:   props,
			results: results,
			count:   0,
		}

		if err := scraper.run(); err != nil {
			log.Printf("scraper error: %v", err)
			return
		}
	}()

	return results
}


// run configures and starts the scraping process.
func (s *scraper) run() error {
	if err := s.setupFileOutput(); err != nil {
		return fmt.Errorf("failed to setup file output: %w", err)
	}
	defer s.closeFile()

	collector := s.setupCollector()

	baseURL, err := url.Parse(s.props.Url)
	if err != nil {
		return fmt.Errorf("failed to parse base URL: %w", err)
	}

	s.configureCollectorHandlers(collector, baseURL)

	if err := collector.Visit(s.props.Url); err != nil {
		return fmt.Errorf("failed to visit initial URL: %w", err)
	}

	return nil
}

// setupCollector initializes a new colly collector.
func (s *scraper) setupCollector() *colly.Collector {
	return colly.NewCollector()
}

// configureCollectorHandlers sets up the HTML and response handlers for the collector.
func (s *scraper) configureCollectorHandlers(c *colly.Collector, baseURL *url.URL) {
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		s.handleLink(e, baseURL)
	})

	c.OnResponse(func(r *colly.Response) {
		s.handleResponse(r)
	})
}

// handleLink processes found links, resolves them, and visits them if they are within the same domain.
func (s *scraper) handleLink(e *colly.HTMLElement, baseURL *url.URL) {
	if s.count >= s.props.Limit {
		return
	}

	href := e.Attr("href")

	if !s.isValidLink(href) {
		return
	}

	absoluteURL := resolveURL(baseURL, href)

	if isSameDomain(baseURL, absoluteURL) {
		cleanURL := stripQueryParams(absoluteURL)
		e.Request.Visit(cleanURL)
	}
}

// handleResponse processes the response, sends the result to the channel, and writes to the output file.
func (s *scraper) handleResponse(r *colly.Response) {
	cleanURL := stripQueryParams(r.Request.URL.String())

	completedURL := NewCompletedUrl(s.count, r.StatusCode, cleanURL)
	s.results <- completedURL

	s.writeToFile(s.count, r.StatusCode, cleanURL)
	s.count++
}
