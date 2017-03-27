package main

import (
    "fmt"
    "os"
    "strconv"
)

func isPrime(n int) bool {
    for i := 2; i <= n / 2; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func main() {
    from, _ := strconv.Atoi(os.Args[1])
    to, _ := strconv.Atoi(os.Args[2])
    for n := from; n <= to; n++ {
        if isPrime(n) {
            fmt.Println(n)
        }
    }
}
