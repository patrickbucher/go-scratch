package prime

func GetPrimesSync(from, to int) []int {
	primes := make([]int, 0)
	for i := from; i <= to; i++ {
		if isPrimeSync(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrimeSync(number int) bool {
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func GetPrimesAsync(from, to int, primes chan<- int) {
	numberChan := make(chan int)
	go isPrimeAsync(numberChan, primes)
	for i := from; i <= to; i++ {
		numberChan <- i
	}
	close(numberChan)
}

func isPrimeAsync(numbers <-chan int, primes chan<- int) {
	for number := range numbers {
		possiblePrime := true
		for i := 2; possiblePrime && i <= number/2; i++ {
			if number%i == 0 {
				possiblePrime = false
			}
		}
		if possiblePrime {
			primes <- number
		}
	}
	close(primes)
}
