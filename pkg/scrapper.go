package scrapper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

var count = 0

func Scrap(props *scrapperProps) *colly.Collector {
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

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())

		if (props.isFile){
			fmt.Fprintln(props.File, r.URL.String())
		}

		count++
	})

	return c
}

func clearUrl(url string) string {
	lastIndex := strings.LastIndex(url, "/")
	return url[:lastIndex+1]
}

func GetCount() int {
	return count
}
