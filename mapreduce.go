package main

import (
	"fmt"
	"strings"
)

func main() {
	mapper := strings.ToUpper
	reducer := func(a, b string) string {
		return a + "|" + b
	}
	values := []string{"this", "is", "a", "test"}
	fmt.Println(Reduce(Map(values, mapper), reducer))
}

func Map(str []string, f func(string) string) []string {
	mapped := make([]string, len(str))
	for k := range str {
		mapped[k] = f(str[k])
	}
	return mapped
}

func Reduce(str []string, f func(string, string) string) string {
	if len(str) < 2 {
		return str[0]
	}
	a := str[0]
	for n := 1; n < len(str); n++ {
		b := str[n]
		a = f(a, b)
	}
	return a
}
