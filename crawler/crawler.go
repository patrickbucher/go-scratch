// Crawler crawls the links of a website.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: crawler [url]\n")
		os.Exit(1)
	}
	url := os.Args[1]
	content, err := getPageContent(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching for url %v: %v\n", url, err)
		os.Exit(1)
	}
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
	now := time.Now()
	subdir := now.Format("2006-01-02-15:04:05")
	dir := strings.Join([]string{"dump", subdir}, string(os.PathSeparator))
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating dir %s: %v\n", dir, err)
		os.Exit(1)
	}
	for k, _ := range ids {
		artUrl := url + "/" + k
		artContent, err := getPageContent(artUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error downloading %s: %v\n", artUrl, err)
		}
		artFile := strings.Replace(k, ";", "", 1)
		artFile = strings.Replace(artFile, ",", "-", 1)
		artFilePath := dir + string(os.PathSeparator) + artFile
		f, err := os.Create(artFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening %s: %v\n", artFilePath, err)
			continue
		}
		f.WriteString(artContent)
		f.Sync()
	}
}

func getPageContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}
