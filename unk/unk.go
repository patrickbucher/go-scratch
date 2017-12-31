// unk.go is a possible solution for a specific problem posted on StackExchange
// https://unix.stackexchange.com/questions/413664/replace-string-in-a-huge-70gb-one-line-text-file
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	const (
		pattern     = "<unk>"
		replacement = "<raw_unk>"
	)
	var match int
	var char rune
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		char = rune(scanner.Text()[0])
		if char == []rune(pattern)[match] {
			match++
			if match == len(pattern) {
				fmt.Print(replacement)
				match = 0
			}
		} else {
			if match > 0 {
				fmt.Print(string(pattern[:match]))
				match = 0
			}
			if char == rune(pattern[0]) {
				match = 1
			} else {
				fmt.Print(string(char))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
