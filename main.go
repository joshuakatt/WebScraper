package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
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

	// Set up the Gin router
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Define the API endpoint for scraping
	router.GET("/api/scrape", handleScrape)

	// Serve the UI
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handleIndex)

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}

	// Wait for the "quit" signal to stop the program
	<-stop
	close(quit)
}

func scrapePage() string {
	// Create a new collector instance
	c := colly.NewCollector()

	// Variable to store the scraped data
	var currentData string

	// Define a callback for monitoring changes
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Extract the current data from the webpage
		currentData = e.Text
	})

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Visit the webpage
	c.Visit("https://www.wikipedia.org/")

	return currentData
}

func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func handleScrape(c *gin.Context) {
	data := scrapePage()
	c.JSON(http.StatusOK, gin.H{"data": data})
}
