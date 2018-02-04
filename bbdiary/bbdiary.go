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
	maxPage       = 244 // TODO: determine by fetching the first page
)

func main() {
	links := make(chan string)
	go collectArticleLinks(links, maxPage)
	for link := range links {
		fmt.Println(link)
	}
}

func collectArticleLinks(links chan<- string, maxPage int) {
	var nPages, nDocs int
	pageDone := make(chan bool)
	documents := make(chan *html.Node)
	for nPages = 0; nPages < maxPage; nPages++ {
		url := fmt.Sprintf(nthPage, nPages)
		go fetchPage(url, documents)
	}
	for nDocs < nPages {
		if doc := <-documents; doc != nil {
			go crawlForHrefs(doc, readLinkClass, links, pageDone)
			nDocs++
		}
	}
	close(documents)
	for i := 0; i < nDocs; i++ {
		<-pageDone
	}
	close(pageDone)
	close(links)
}

func fetchPage(url string, document chan<- *html.Node) {
	var statusCode int
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get %s failed: %v", url, err)
		document <- nil
	}
	defer resp.Body.Close()
	if statusCode = resp.StatusCode; statusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Get %s failed: %s", url, resp.Status)
		document <- nil
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %s: %v", url, err)
		document <- nil
	}
	document <- doc
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
