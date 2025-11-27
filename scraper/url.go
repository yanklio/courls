package scraper

import (
	"net/url"
	"strings"
)

// resolveURL resolves a URL relative to a base URL and removes query parameters.
func resolveURL(base *url.URL, href string) string {
	u, err := url.Parse(href)
	if err != nil {
		return href
	}
	resolved := base.ResolveReference(u)
	resolved.RawQuery = ""
	return resolved.String()
}

// stripQueryParams removes the query parameters from a URL.
func stripQueryParams(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}
	u.RawQuery = ""
	return u.String()
}

// isSameDomain checks if a given URL is in the same domain as the base URL.
func isSameDomain(base *url.URL, hrefStr string) bool {
	u, err := url.Parse(hrefStr)
	if err != nil {
		return false
	}
	return u.Host == base.Host
}

// isValidLink checks if a given link is valid for scraping.
func (s *scraper) isValidLink(href string) bool {
	if href == "" {
		return false
	}

	if strings.Contains(href, "#") {
		return false
	}

	if strings.Contains(href, ".") {
		return false
	}

	return true
}
