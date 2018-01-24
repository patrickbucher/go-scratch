package sorting

func BubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if numbers[j-1] > numbers[j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}
		}
	}
}
