// Crawler crawls the Articles of a website and downloads them.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s [url]", os.Args[0])
	}
	url := os.Args[1]
	content, err := getPageContent(url)
	if err != nil {
		log.Fatalf("error fetching for url %v: %v", url, err)
	}
	now := time.Now()
	subdir := now.Format("2006-01-02_15.04.05")
	dir := strings.Join([]string{"dump", subdir}, string(os.PathSeparator))
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating dir %s: %v\n", dir, err)
		os.Exit(1)
	}
	ids := extractArticleIds(content)
	var n int
	ch := make(chan bool)
	for _, bar := range ids {
		if len(bar) == 0 {
			continue
		}
		artURL := url + "/" + bar
		artFile := strings.Replace(bar, ";", "", 1)
		artFile = strings.Replace(artFile, ",", "-", 1)
		artFilePath := dir + string(os.PathSeparator) + artFile
		go downloadArticle(artURL, artFilePath, ch)
		n++
	}
	var good, bad int
	for i := 0; i < n; i++ {
		success := <-ch
		if success {
			good++
		} else {
			bad++
		}
	}
	fmt.Printf("downloaded: %d, failed: %d\n", good, bad)
}

func downloadArticle(url, path string, done chan bool) {
	start := time.Now()
	artContent, err := getPageContent(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error downloading %s: %v\n", url, err)
	}
	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening %s: %v\n", path, err)
		done <- false
	}
	f.WriteString(artContent)
	f.Sync()
	defer f.Close()
	elapsed := time.Since(start)
	fmt.Printf("downloaded '%s' to '%s' in %v\n", url, path, elapsed)
	done <- true
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

func extractArticleIds(content string) []string {
	r, err := regexp.Compile(`href="[^;]*(;art[^"]+)"`)
	if err != nil {
		log.Fatalf("error parsing regex: %v", err)
	}
	hrefs := r.FindAllStringSubmatch(content, -1)
	ids := make(map[string]struct{})
	for _, h := range hrefs {
		id := h[1]
		_, contained := ids[id]
		if !contained {
			ids[id] = struct{}{}
		}
	}
	uniqueIds := make([]string, len(ids))
	for k, _ := range ids {
		uniqueIds = append(uniqueIds, k)
	}
	return uniqueIds
}
