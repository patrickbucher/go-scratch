// unk.go is a possible solution for a specific problem posted on StackExchange
// https://unix.stackexchange.com/questions/413664/replace-string-in-a-huge-70gb-one-line-text-file
package unk

import (
	"bufio"
	"io"
)

const (
	pattern     = "<unk>"
	replacement = "<raw_unk>"
)

func ReplaceUnk(input io.Reader, output io.Writer) error {
	var match int
	var char byte
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		char = scanner.Text()[0]
		if char == []byte(pattern)[match] {
			match++
			if match == len(pattern) {
				output.Write([]byte(replacement))
				match = 0
			}
		} else {
			if match > 0 {
				output.Write([]byte(pattern[:match]))
				match = 0
			}
			if char == pattern[0] {
				match = 1
			} else {
				output.Write([]byte{char})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
