package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// Parse the command-line arguments
	if len(os.Args) != 2 {
		fmt.Println("USAGE: go run main.go URL")
		return
	}
	url := os.Args[1]

	// Make an HTTP GET Request
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making an HTTP Request: %s", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body: %s", err)
		return
	}

	// Parse the HTML document
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		fmt.Printf("Error parsing HTML: %s", err)
		return
	}

	// traversing the document and printing data and attributes of all
	// elements
	var TraverseAndPrint func(*html.Node)
	TraverseAndPrint = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("Element: %s\n", n.Data)
			for _, attr := range n.Attr {
				fmt.Printf("  Attribute: %s=%s\n", attr.Key, attr.Val)
			}
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				fmt.Printf("  Data: %s\n", n.FirstChild.Data)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			TraverseAndPrint(c)
		}
	}
	TraverseAndPrint(doc)
}
