package fib

func Fib(n int) int {
	if n <= 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

var cache = make(map[int]int)

func FibCached(n int) int {
	if len(cache) == 0 {
		cache[1] = 1
		cache[2] = 1
	}
	if val, ok := cache[n]; ok {
		return val
	}
	fib := FibCached(n-1) + FibCached(n-2)
	cache[n] = fib
	return fib
}
