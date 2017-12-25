// Crawler crawls the links of a website.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: crawler [url]\n")
		os.Exit(1)
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching url %v: %v\n", url, err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading content: %v\n", err)
	}
	content := string(body)
	r, _ := regexp.Compile(`href="[^;]*(;art[^"]+)"`)
	hrefs := r.FindAllStringSubmatch(content, -1)
	ids := make(map[string]struct{})
	for _, h := range hrefs {
		id := h[1]
		_, contained := ids[id]
		if !contained {
			ids[id] = struct{}{}
		}
	}
	for k, _ := range ids {
		fmt.Println(k)
	}
}
