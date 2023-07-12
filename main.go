package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gocolly/colly/v2"
)

var (
	previousData string
	quit         = make(chan struct{})
)

func main() {
	// Start the background goroutine to periodically scrape the page
	ticker := time.NewTicker(1 * time.Minute) // Adjust the interval as per your requirement
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		for {
			select {
			case <-ticker.C:
				scrapePage()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Wait for the "quit" signal to stop the program
	<-stop
	close(quit)
}

func scrapePage() {
	// Create a new collector instance
	c := colly.NewCollector()

	// Define a callback for monitoring changes
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Extract the current data from the webpage
		currentData := e.Text

		// Compare the current data with the previous data
		if currentData != previousData {
			// If there's an update, inform the user in the terminal
			fmt.Println("There's an update on the monitored website!")

			// Update the previous data with the current data for future comparisons
			previousData = currentData
		}
	})

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Visit the webpage
	c.Visit("https://www.wikipedia.org/")
}
