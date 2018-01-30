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
	links := collectArticleLinks()
	for _, v := range links {
		fmt.Println(v)
	}
}

func collectArticleLinks() []string {
	var links []string
	var statusCode int
	for page := 1; statusCode != 404; page++ {
		url := fmt.Sprintf(nthPage, page)
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
		hrefs := crawlForHrefs(doc, readLinkClass)
		for _, v := range hrefs {
			links = append(links, v)
		}
	}
	return links
}

func crawlForHrefs(n *html.Node, class string) []string {
	var hrefs []string
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
			hrefs = append(hrefs, href)
			fmt.Println(href)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		for _, v := range crawlForHrefs(c, class) {
			hrefs = append(hrefs, v)
		}
	}
	return hrefs
}
