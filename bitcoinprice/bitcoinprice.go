package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const url = "http://blockchain.info/de/ticker"

type price struct {
	Symbol string  `json:"symbol"`
	Price  float32 `json:"last"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s currencies\n", os.Args[0])
		os.Exit(1)
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get %s: %v\n", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "get %s: %s\n", url, resp.Status)
		os.Exit(1)
	}
	buf := bytes.NewBufferString("")
	io.Copy(buf, resp.Body)
	var entries map[string]price
	err = json.Unmarshal(buf.Bytes(), &entries)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal JSON response: %q %v\n", string(buf.Bytes()), err)
		os.Exit(1)
	}
	for _, arg := range os.Args[1:] {
		cur := strings.ToUpper(arg)
		if p, ok := entries[cur]; ok {
			fmt.Printf("%s:%10.2f %s\n", cur, p.Price, p.Symbol)
		} else {
			fmt.Fprintf(os.Stderr, "currency '%s' not in response\n", cur)
		}
	}
}
