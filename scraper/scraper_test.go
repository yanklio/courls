package scraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestScrap(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.URL.Path == "/" {
			fmt.Fprintln(w, `<a href="/page2">Page 2</a>`)
		} else if r.URL.Path == "/page2" {
			fmt.Fprintln(w, `<a href="/">Home</a>`)
		} else {
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	props := NewScraperProps(server.URL, 5, "") // Increased limit to ensure it can crawl both pages
	results := Scrap(props)

	var completed []*CompletedUrl
	timeout := time.After(5 * time.Second)

	// Consume results from the channel
	func() {
		for {
			select {
			case result, ok := <-results:
				if !ok {
					return // Channel closed
				}
				completed = append(completed, result)
			case <-timeout:
				t.Fatal("test timed out waiting for results")
				return
			}
		}
	}()

	// --- Assertions ---
	if len(completed) < 2 {
		t.Fatalf("expected at least 2 results, got %d", len(completed))
	}

	// Extract URLs for easier comparison
	scrapedURLs := make(map[string]struct{})
	for _, c := range completed {
		// Normalize by removing trailing slash for consistent comparison
		normalizedURL := strings.TrimSuffix(c.Url, "/")
		scrapedURLs[normalizedURL] = struct{}{}
	}

	expectedURLs := []string{
		server.URL,
		server.URL + "/page2",
	}

	for _, expectedURL := range expectedURLs {
		u, _ := url.Parse(expectedURL)
		cleanURL := stripQueryParams(u.String())
		// Also normalize the expected URL
		normalizedExpectedURL := strings.TrimSuffix(cleanURL, "/")
		if _, ok := scrapedURLs[normalizedExpectedURL]; !ok {
			// For better error reporting
			var urls []string
			for k := range scrapedURLs {
				urls = append(urls, k)
			}
			t.Errorf("expected URL %s was not scraped. Scraped URLs: %v", normalizedExpectedURL, urls)
		}
	}
}
