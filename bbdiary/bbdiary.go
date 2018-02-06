package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const (
	indexPage       = "https://bonnerandpartners.com/category/dre/"
	nthPage         = "https://bonnerandpartners.com/category/dre/page/%d/"
	readLinkClass   = "w-blog-post-more"
	pageNumberClass = "page-numbers"
)

func main() {
	maxPage, err := determineMaxPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "determine page: %v\n", err)
		os.Exit(1)
	}
	links := make(chan string)
	go collectArticleLinks(links, maxPage)
	for link := range links {
		fmt.Println(link)
	}
}

func determineMaxPage() (int, error) {
	resp, err := http.Get(indexPage)
	if err != nil {
		return 0, fmt.Errorf("get page %s: %v\n", indexPage, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("get page %s: %s", indexPage, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("parsing document: %v\n", err)
	}
	var max int
	if spans := crawlForFirstChild(doc, "a", pageNumberClass); len(spans) > 0 {
		for _, span := range spans {
			if span.FirstChild != nil && span.FirstChild.Type == html.TextNode {
				if n, err := strconv.Atoi(span.FirstChild.Data); err == nil {
					if n > max {
						max = n
					}
				}
			}
		}
		return max, nil
	}
	return 0, fmt.Errorf("can't determine max page number")
}

func crawlForFirstChild(n *html.Node, element, class string) []*html.Node {
	var children []*html.Node
	if n.Type == html.ElementNode && n.Data == element {
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == class {
				children = append(children, n.FirstChild)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		for _, v := range crawlForFirstChild(c, element, class) {
			children = append(children, v)
		}
	}
	return children
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
