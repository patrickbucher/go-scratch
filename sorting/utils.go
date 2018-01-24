package sorting

import (
	"math/rand"
	"time"
)

func RandomIntSlice(n, max int) []int {
	rand.Seed(time.Now().UnixNano())
	if n < 1 || max < 1 {
		return []int{}
	}
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Int() % (max + 1)
	}
	return numbers
}

func CheckSorted(numbers []int) bool {
	l := len(numbers)
	for i := 1; i < l; i++ {
		if numbers[i-1] > numbers[i] {
			return false
		}
	}
	return true
}
