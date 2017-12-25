package main

import (
	"accent"
	"bufio"
	"fmt"
	"io"
	"os"
)

func getAccentFromArgs() rune {
	if len(os.Args) >= 2 {
		param := os.Args[1]
		paramRunes := []rune(param)
		return paramRunes[0]
	} else {
		// default accent
		return '*'
	}
}

func main() {
	pseudoAccent := getAccentFromArgs()
	reader := bufio.NewReader(os.Stdin)
	eof := false
	for !eof {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				eof = true
			} else {
				panic(fmt.Sprintf("%v", err))
			}
		}
		fmt.Print(accent.Accent(line, pseudoAccent))
	}
}
