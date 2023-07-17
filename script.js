// script.js

document.getElementById("scrapeForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const url = document.getElementById("urlInput").value;
  
    // Send the URL to the server-side code for scraping
    fetch('/api/scrape', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ url }),
    })
    .then(response => response.json())
    .then(data => {
        // Update the UI with the scraped data
        console.log(data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
  });