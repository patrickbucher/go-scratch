// Crawler crawls the Articles of a website and downloads them.
package main

import (
	"bytes"
	"fmt"
	"io"
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
	buf := bytes.NewBufferString("")
	reader, err := getPageContent(url)
	if err != nil {
		log.Fatalf("error fetching for url %v: %v", url, err)
	}
	defer reader.Close()
	buf.ReadFrom(reader)
	content := buf.String()
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
	var succeeded, failed int
	for i := 0; i < n; i++ {
		success := <-ch
		if success {
			succeeded++
		} else {
			failed++
		}
	}
	close(ch)
	fmt.Printf("downloaded: %d, failed: %d\n", succeeded, failed)
}

func downloadArticle(url, path string, done chan bool) {
	start := time.Now()
	reader, err := getPageContent(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error downloading %s: %v\n", url, err)
	}
	defer reader.Close()
	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening %s: %v\n", path, err)
		done <- false
	}
	defer f.Close()
	io.Copy(f, reader)
	elapsed := time.Since(start)
	fmt.Printf("downloaded '%s' to '%s' in %v\n", url, path, elapsed)
	done <- true
}

func getPageContent(url string) (io.ReadCloser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GET request to %s: %v", url, err)
	}
	req.Header.Set("User-Agent", "Googlebot-News")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET request %v caused error: %v", req, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET request %v failed: %s", url, resp.Status)
	}
	return resp.Body, nil
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
	for k := range ids {
		uniqueIds = append(uniqueIds, k)
	}
	return uniqueIds
}
