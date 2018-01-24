package sorting

import "testing"

const (
	testSize  = 10000   // numbers of items to be sorted
	maxNumber = 1000000 // biggest item possible
)

func TestBubbleSort(t *testing.T) {
	numbers := RandomIntSlice(testSize, maxNumber)
	BubbleSort(numbers)
	if ok := CheckSorted(numbers); !ok {
		t.Fatalf("%v is not sorted", numbers)
	}
}

func TestBubbleSortOptimized(t *testing.T) {
	numbers := RandomIntSlice(testSize, maxNumber)
	BubbleSortOptimized(numbers)
	if ok := CheckSorted(numbers); !ok {
		t.Fatalf("%v is not sorted", numbers)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	numbers := RandomIntSlice(testSize, maxNumber)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(numbers)
	}
	b.StopTimer()
}

func BenchmarkBubbleSortOptimized(b *testing.B) {
	numbers := RandomIntSlice(testSize, maxNumber)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		BubbleSortOptimized(numbers)
	}
	b.StopTimer()
}
