package sum

func Sum(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

func SumGauss(n int) int {
    return n * (n + 1) / 2
}
