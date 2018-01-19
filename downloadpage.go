package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s [url]", os.Args[0])
	}
	r, err := getPageContent(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	io.Copy(os.Stdout, r)
}

func getPageContent(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("GET %s failed, cause: %s", url, res.Status)
		return nil, errors.New(msg)
	}
	return res.Body, nil
}
