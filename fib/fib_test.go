package fib

import "testing"

var tests = []struct {
	n    int
	want int
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
}

func TestFib(t *testing.T) {
	for _, test := range tests {
		got := Fib(test.n)
		if got != test.want {
			t.Errorf("Fib(%d) got: %d, want: %d", test.n, got, test.want)
		}
	}
}

func TestFibCached(t *testing.T) {
	for _, test := range tests {
		got := FibCached(test.n)
		if got != test.want {
			t.Errorf("FibCached(%d) got: %d, want: %d", test.n, got, test.want)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(40)
	}
}

func BenchmarkFibCached(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibCached(10)
	}
}
