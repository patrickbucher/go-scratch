package main

import (
	"fmt"
)

func write(ch chan<- int) {
	for n := 1; n <= 10; n++ {
		ch <- n
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go write(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
