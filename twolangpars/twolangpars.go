// twolangpars: adds two language attributes alternately to a tag
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		first  = "ru"
		second = "de"
		tag    = "<p>"
		tagfmt = `<p lang="%s">`
	)
	if len(os.Args[1:]) == 0 {
		fmt.Fprintf(os.Stderr, "usage: %s file_1 [file_2 ...]\n", os.Args[0])
		return
	}
	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to read file %s\n", file)
			continue
		}
		input := bufio.NewScanner(f)
		text := ""
		for input.Scan() {
			line := input.Text()
			text += line + "\n"
		}
		tokens := strings.Split(text, tag)
		fmt.Print(tokens[0])
		for i, t := range tokens[1:len(tokens)] {
			var lang string
			if i%2 == 0 {
				lang = first
			} else {
				lang = second
			}
			langtag := fmt.Sprintf(tagfmt, lang)
			fmt.Print(langtag, t)
		}
	}
}
