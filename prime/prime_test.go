package prime

import "testing"

func TestGetPrimesSync(t *testing.T) {
    primes := GetPrimesSync(10, 20)
    if primes[0] != 11 {
        t.Error(`GetPrimesSync(10, 20): 11 not found`)
    }
    if primes[1] != 13 {
        t.Error(`GetPrimesSync(10, 20): 13 not found`)
    }
    if primes[2] != 17 {
        t.Error(`GetPrimesSync(10, 20): 17 not found`)
    }
}

func TestGetPrimesAsync(t *testing.T) {
    primes := make(chan int)
    go GetPrimesAsync(10, 20, primes)
    if <-primes != 11 {
        t.Error(`GetPrimesSync(10, 20): 11 not found`)
    }
    if <-primes != 13 {
        t.Error(`GetPrimesSync(10, 20): 13 not found`)
    }
    if <-primes != 17 {
        t.Error(`GetPrimesSync(10, 20): 17 not found`)
    }
}

func BenchmarkGetPrimesSync(b *testing.B) {
    for i := 0; i <= b.N; i++ {
        GetPrimesSync(1, 100)
    }
}

func BenchmarkGetPrimesAsync(b *testing.B) {
    for i := 0; i <= b.N; i++ {
        primes := make(chan int)
        go GetPrimesAsync(1, 100, primes)
    }
}
