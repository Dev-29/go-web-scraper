# Golang Web Scraper

This is a CLI-based web scraper written in Go. It makes an HTTP GET request to a specified URL, parses the HTML document, and traverses the document to print the data and attributes of all elements.

## Dependencies

- Go 1.16 or higher
- The `html` package from the `golang.org/x/net` package

## Building and Running

To build and run the program, you will need to install the dependencies and then run the following commands:

`go build main.go`

`./main https://www.example.com`

Replace `https://www.example.com` with the URL of the website you want to scrape.
