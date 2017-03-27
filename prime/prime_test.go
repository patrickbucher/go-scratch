package prime

import "testing"

var tests = []struct {
    from int
    to int
    want []int
}{
    {10, 20, []int{11, 13, 17, 19}},
    {20, 30, []int{23, 29}},
    {95, 99, []int{97}},
}

func TestGetPrimesSync(t *testing.T) {
    for _, test := range tests {
        got := GetPrimesSync(test.from, test.to)
        if !eqList(got, test.want) {
            t.Errorf("GetPrimeSync(%d, %d) got: %d, want: %d",
                test.from, test.to, got, test.want)
        }
    }
}

func TestGetPrimesAsync(t *testing.T) {
    for _, test := range tests {
        primes := make(chan int)
        go GetPrimesAsync(test.from, test.to, primes)
        got := make([]int, 0)
        for p := range primes {
            got = append(got, p)
        }
        if !eqList(got, test.want) {
            t.Errorf("GetPrimeAsync(%d, %d) got: %d, want: %d",
                test.from, test.to, got, test.want)
        }
    }
}

func eqList(a, b []int) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil && b != nil {
        return false
    }
    if a != nil && b == nil {
        return false
    }
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func BenchmarkGetPrimesSync(b *testing.B) {
    for i := 0; i <= b.N; i++ {
        GetPrimesSync(1, 1000)
    }
}

func BenchmarkGetPrimesAsync(b *testing.B) {
    for i := 0; i <= b.N; i++ {
        primes := make(chan int)
        go GetPrimesAsync(1, 1000, primes)
    }
}
