package scrapper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)


func Scrap(props *scrapperProps) *scrapperResult {
	count := 0
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
		fmt.Printf("%d\t%d\t%s\n", count, r.StatusCode, r.Request.URL.String())

		if (props.isFile){
			fmt.Fprintf(props.File, "%d\t%d\t%s\n", count, r.StatusCode, r.Request.URL.String())
		}

		count++
	})

	fmt.Println("Count   Code    URL")
	fmt.Println("------  -----   ------------------------")

	c.Visit(props.Url)

	res := NewScrapperResult(count)

	return res
}

func clearUrl(url string) string {
	lastIndex := strings.LastIndex(url, "/")
	return url[:lastIndex+1]
}
