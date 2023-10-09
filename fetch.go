package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/net/html"
)

func main() {
	metadataPtr := flag.Bool("metadata", false, "Record metadata about what was fetched")
	flag.Parse()

	for _, s := range flag.Args() {
		// Validate URL
		u, err := url.Parse(s)
		if err != nil {
			fmt.Println("Invalid url", s)
			continue
		}

		// Fetch URL
		response, err := http.Get(s)
		if err != nil {
			fmt.Println("Error fetching", s)
			continue
		}
		defer response.Body.Close()

		// Read response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %s\n", err)
			continue
		}

		// Metadata
		if *metadataPtr {
			links, images := countNumberOfLinksAndImages(response.Body)
			fmt.Printf("site: %s\n", u.Host)
			fmt.Printf("num_links: %d\n", links)
			fmt.Printf("images: %d\n", images)
			fmt.Printf("last_fetch: %s\n", time.Now().UTC().Format("Mon Jan 01 2006 15:04 UTC"))
		}

		// Write to file
		filename := u.Host + ".html"
		err = os.WriteFile(filename, body, 0644)
		if err != nil {
			fmt.Printf("Error writing to file: %s\n", err)
			continue
		}
	}
}

func countNumberOfLinksAndImages(body io.Reader) (int, int) {
	links := 0
	images := 0
	tokenizer := html.NewTokenizer(body)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return links, images
		} else if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {
			token := tokenizer.Token()
			if token.Data == "a" {
				links++
			}
			if token.Data == "img" {
				images++
			}
		}
	}
}
