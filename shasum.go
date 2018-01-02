package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
)

func main() {
	input := "This is a test."
	sha256, _ := shaSum(input, 256)
	sha384, _ := shaSum(input, 384)
	sha512, _ := shaSum(input, 512)
	fmt.Println(input, sha256, sha384, sha512)
}

func shaSum(data string, size uint) ([]byte, error) {
	input := []byte(data)
	switch size {
	case 256:
		bytes := sha256.Sum256(input)
		return bytes[:], nil
	case 384:
		bytes := sha512.Sum384(input)
		return bytes[:], nil
	case 512:
		bytes := sha512.Sum512(input)
		return bytes[:], nil
	default:
		return nil, errors.New("unsupported sha size")
	}
}
