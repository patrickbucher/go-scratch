package main

import (
    "accent"
    "bufio"
    "fmt"
    "io"
    "os"
)

func getAccentFromArgs() rune {
    if len(os.Args) < 2 {
        panic("usage: accent [pseudo accent sign]")
    }
    param := os.Args[1]
    paramRunes := []rune(param)
    if len(paramRunes) > 1 {
        panic("you must only one character as pseudo accent sign")
    }
    return paramRunes[0]
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
