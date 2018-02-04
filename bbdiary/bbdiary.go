package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const (
	nthPage       = "https://bonnerandpartners.com/category/dre/page/%d/"
	readLinkClass = "w-blog-post-more"
)

func main() {
	links := make(chan string)
	go collectArticleLinks(links)
	for link := range links {
		fmt.Println(link)
	}
}

func collectArticleLinks(links chan<- string) {
	var statusCode, page int
	pageDone := make(chan bool)
	for page = 1; statusCode != 404; page++ {
		url := fmt.Sprintf(nthPage, page)
		// TODO: This part still works synchronously. Retrieve documents in a
		// fire-and-forget manner by sending the resulting DocumentNode through
		// a channel.
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Get %s failed: %v", url, err)
		}
		defer resp.Body.Close()
		if statusCode = resp.StatusCode; statusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "Get %s failed: %s", url, resp.Status)
			continue
		}
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parsing %s: %v", url, err)
			continue
		}
		go crawlForHrefs(doc, readLinkClass, links, pageDone)
	}
	for i := 0; i < page; i++ {
		<-pageDone
	}
	close(pageDone)
	close(links)
}

func crawlForHrefs(n *html.Node, class string, links chan<- string, done chan<- bool) {
	if n.Type == html.ElementNode {
		href := ""
		found := false
		for _, a := range n.Attr {
			if a.Key == "href" {
				href = a.Val
			} else if a.Key == "class" && strings.Contains(a.Val, class) {
				found = true
			}
		}
		if found {
			links <- href
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		go crawlForHrefs(c, class, links, done)
	}
	if n.Type == html.DocumentNode {
		done <- true
	}
}
