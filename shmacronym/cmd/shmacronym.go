package main

import (
	"fmt"
	"math/rand"
	"os"
	"shmacronym"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [acronym]", os.Args[0])
		return
	}
	acronym := os.Args[1]
	n := len(acronym)
	for i := 0; i < n; i++ {
		r := rune(acronym[i])
		var word string
		var remainder []string
		if i != n-1 {
			word, remainder = extract(shmacronym.Adjectives[r])
			shmacronym.Adjectives[r] = remainder
		} else {
			word, remainder = extract(shmacronym.Nouns[r])
			shmacronym.Nouns[r] = remainder
		}
		if word != "" {
			fmt.Println(word)
		} else {
			fmt.Fprintf(os.Stderr, "no more words left for '%c'\n", r)
		}
	}
}

func extract(words []string) (string, []string) {
	if len(words) == 0 {
		return "", make([]string, 0)
	}
	var word string
	var remainder []string
	x := rand.Intn(len(words))
	word = words[x]
	remainder = append(words[:x], words[x+1:]...)
	return word, remainder
}
