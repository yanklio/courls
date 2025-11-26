package scraper

import (
	"net/url"
	"strings"
)

func resolveURL(base *url.URL, href string) string {
	u, err := url.Parse(href)
	if err != nil {
		return href
	}
	resolved := base.ResolveReference(u)
	resolved.RawQuery = ""
	return resolved.String()
}

func stripQueryParams(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}
	u.RawQuery = ""
	return u.String()
}

func isSameDomain(base *url.URL, hrefStr string) bool {
	u, err := url.Parse(hrefStr)
	if err != nil {
		return false
	}
	return u.Host == base.Host
}

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
