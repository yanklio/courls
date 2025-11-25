package scrapper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)


func Scrap(props *scrapperProps) <-chan *CompletedUrl{

	count := 0
	results := make(chan *CompletedUrl)

	go func() {
		defer close(results)

		c := colly.NewCollector()
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			if count >= props.Limit {
				return
			}

			href := (e.Attr("href"))

			if strings.HasPrefix(href, props.Url) {
				e.Request.Visit(clearUrl(href))
			}
		})

		c.OnResponse(func(r *colly.Response) {
			results <- NewCompletedUrl(count, r.StatusCode, r.Request.URL.String())

			count++
		})

		c.Visit(props.Url)
	}()

	return results
}

func clearUrl(url string) string {
	lastIndex := strings.LastIndex(url, "/")
	return url[:lastIndex+1]
}
