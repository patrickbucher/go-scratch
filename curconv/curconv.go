package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const api = "https://api.fixer.io/latest?base=%s"

type apiResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func processArgs() (amount float64, from, to string) {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "usage: %s [amount] [from] [to]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "example: %s 100 EUR USD\n", os.Args[0])
		os.Exit(1)
	}
	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse '%s' to float64: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	from, to = os.Args[2], os.Args[3]
	if len(from) != 3 || len(to) != 3 {
		fmt.Fprintf(os.Stderr, "currency codes must be three characters long\n")
		os.Exit(1)
	}
	return amount, strings.ToUpper(from), strings.ToUpper(to)
}

func outputResult(fromCur, toCur string, fromVal, toVal float64) {
	fromLen := len(fmt.Sprintf("%.2f", fromVal))
	toLen := len(fmt.Sprintf("%.2f", toVal))
	var len int
	if fromLen > toLen {
		len = fromLen
	} else {
		len = toLen
	}
	format := "%s %" + strconv.Itoa(len) + ".2f\n"
	fmt.Printf(format+format, fromCur, fromVal, toCur, toVal)
}

func main() {
	amount, from, to := processArgs()
	url := fmt.Sprintf(api, from)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GET %s: %v\n", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "GET %s: %s\n", url, resp.Status)
		os.Exit(1)
	}
	buf := bytes.NewBufferString("")
	io.Copy(buf, resp.Body)
	var result apiResponse
	err = json.Unmarshal(buf.Bytes(), &result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal %s: %v\n", buf.String(), err)
		os.Exit(1)
	}
	if rate, ok := result.Rates[to]; ok {
		result := amount * rate
		outputResult(from, to, amount, result)
	} else {
		fmt.Fprintf(os.Stderr, "currency '%s' not in result\n", to)
		os.Exit(1)
	}
}
