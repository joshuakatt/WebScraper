package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// Callback to extract text
	c.OnHTML("body", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Println("Text:", text)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://www.donaldjtrump.com/")
}
