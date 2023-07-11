package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/smtp"
)

func main() {
	// Email configuration
	smtpHost := "smtp.example.com"
	smtpPort := 587
	smtpUsername := "your-email@example.com"
	smtpPassword := "your-password"
	sender := "your-email@example.com"
	recipient := "recipient@example.com"

	// Create a new collector instance
	c := colly.NewCollector()

	// Define a callback for monitoring changes
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Perform comparison with previous data to detect updates
		// If there's an update, send an email notification
		// You can customize the comparison logic based on your specific requirements
		// Here's a simple example that always sends an email when the callback is triggered
		sendEmailNotification(smtpHost, smtpPort, smtpUsername, smtpPassword, sender, recipient)
	})

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start the scraping process
	c.Visit("https://example.com")
}

func sendEmailNotification(smtpHost string, smtpPort int, smtpUsername string, smtpPassword string, sender string, recipient string) {
	// Compose the email content
	subject := "Web Scraper Update Notification"
	body := "There's an update on the monitored website!"

	// Set up the SMTP client
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Compose the email message
	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email
	err := smtp.SendMail(smtpHost+":"+fmt.Sprint(smtpPort), auth, sender, []string{recipient}, message)
	if err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Email notification sent successfully")
	}
}
