package scrapper

import (
	"fmt"
	"net/url"
	"os"

	"strings"

	"github.com/gocolly/colly/v2"
)

func Scrap(props *scrapperProps) <-chan *CompletedUrl {
	count := 0

	var file *os.File = nil
	var err error = nil

	results := make(chan *CompletedUrl)

	go func(props *scrapperProps) {
		defer close(results)

		if props.isFile {
			file, err = os.Create(props.FileName)
			if err != nil {
				panic(err)
			}
			defer file.Close()
		}

		c := colly.NewCollector()
		baseURL, err := url.Parse(props.Url)
		if err != nil {
			return
		}

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			if count >= props.Limit {
				return
			}

			href := e.Attr("href")

			if href == "" || strings.Contains(href, "#") || strings.Contains(href, ".") {
				return
			}

			absoluteURL := resolveURL(baseURL, href)

			if isSameDomain(baseURL, absoluteURL) {
				cleanURL := stripQueryParams(absoluteURL)
				e.Request.Visit(cleanURL)
			}
		})

		c.OnResponse(func(r *colly.Response) {
			cleanURL := stripQueryParams(r.Request.URL.String())
			results <- NewCompletedUrl(count, r.StatusCode, cleanURL)
			if (props.isFile) {
				fmt.Fprintf(file, "%5d   %3d     %s\n", count, r.StatusCode, cleanURL)
			}
			count++
		})

		c.Visit(props.Url)
	}(props)

	return results
}

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
