package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(naturals <-chan int, primes chan<- int) {
	for n := range naturals {
		dividable := false
		for i := 2; !dividable && i <= n/2; i++ {
			if n%i == 0 {
				dividable = true
			}
		}
		if !dividable {
			primes <- n
		}
	}
	close(primes)
}

func main() {
	from, _ := strconv.Atoi(os.Args[1])
	to, _ := strconv.Atoi(os.Args[2])
	naturals := make(chan int)
	primes := make(chan int)
	go isPrime(naturals, primes)
	go func() {
		for n := from; n <= to; n++ {
			naturals <- n
		}
		close(naturals)
	}()
	for p := range primes {
		fmt.Println(p)
	}
}
