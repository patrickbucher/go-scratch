package sum

import "testing"

var tests = []struct {
	n    int
	want int
}{
	{10, 55},
	{100, 5050},
	{1000, 500500},
	{10000, 50005000},
}

func TestSum(t *testing.T) {
	for _, test := range tests {
		got := Sum(test.n)
		if got != test.want {
			t.Errorf("Sum(%d): got %d, want %d", test.n, got, test.want)
		}
	}
}

func TestSumGauss(t *testing.T) {
	for _, test := range tests {
		got := SumGauss(test.n)
		if got != test.want {
			t.Errorf("SumGauss(%d): got %d, want %d", test.n, got, test.want)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(100000)
	}
}

func BenchmarkSumGauss(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumGauss(100000)
	}
}
