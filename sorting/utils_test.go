package sorting

import "testing"

func TestRandomIntSlice(t *testing.T) {
	const max = 100
	const n = 1000
	numbers := RandomIntSlice(n, max)
	if l := len(numbers); l != n {
		t.Fatalf("wrong length of numbers: expected %d, got %d\n", max, l)
	}
	for _, v := range numbers {
		if v < 0 || v > max {
			t.Fatalf("wrong number: expected in range [0;%d], got %d\n", max, v)
		}
	}
}

func TestCheckSorted(t *testing.T) {
	tests := []struct {
		numbers []int
		sorted  bool
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, true},
		{[]int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8}, false},
	}
	for _, test := range tests {
		if sorted := CheckSorted(test.numbers); sorted != test.sorted {
			t.Fatalf("is %v sorted? expected: %v, got: %v\n",
				test.numbers, test.sorted, sorted)
		}
	}
}
