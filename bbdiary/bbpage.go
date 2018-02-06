package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const (
	articleURL      = "https://bonnerandpartners.com/cryptocurrencies-your-best-defense-in-the-war-on-cash/"
	articleDivClass = "l-main"
)

func main() {
	resp, err := http.Get(articleURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getting url %s: %v\n", articleURL, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "getting url %s: %s\n", articleURL, resp.Status)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing HTML: %v\n", err)
		os.Exit(1)
	}
	contentNodes := extractArticleContent(doc)
	var textNodes []*html.Node
	for _, n := range contentNodes {
		for _, t := range extractTextNodes(n) {
			textNodes = append(textNodes, t)
			if str := strings.TrimSpace(t.Data); str != "" {
				fmt.Println(str)
			}
		}
	}
}

func extractTextNodes(n *html.Node) []*html.Node {
	var textNodes []*html.Node
	if n.Type == html.TextNode {
		textNodes = append(textNodes, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "style" || c.Data == "script") {
			continue
		}
		for _, t := range extractTextNodes(c) {
			textNodes = append(textNodes, t)
		}
	}
	return textNodes
}

func extractArticleContent(n *html.Node) []*html.Node {
	var articleContent []*html.Node
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == articleDivClass {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode || c.Type == html.ElementNode {
						articleContent = append(articleContent, c)
					}
				}
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		for _, v := range extractArticleContent(c) {
			articleContent = append(articleContent, v)
		}
	}
	return articleContent
}
