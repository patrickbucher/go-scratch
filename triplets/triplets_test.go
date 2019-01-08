package triplets

import (
	"testing"
)

var tests = []struct {
	n        int
	expected []Triangle
}{
	{n: 1000, expected: []Triangle{{a: 200, b: 375, c: 425}}},
}

func BenchmarkTriplets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Triplets(test.n)
		}
	}
}
func BenchmarkTripletsOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			TripletsOptimized(test.n)
		}
	}
}
