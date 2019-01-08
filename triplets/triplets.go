package triplets

type Triangle struct {
	a int
	b int
	c int
}

func Triplets(n int) []Triangle {
	var a, b, c int
	triplets := make([]Triangle, 0)

	for a = 1; a < n; a++ {
		for b = 1; b < n; b++ {
			c = n - a - b
			if c*c == a*a+b*b {
				triplets = append(triplets, Triangle{a: a, b: b, c: c})
			}
		}
	}

	return triplets
}

func TripletsOptimized(n int) []Triangle {
	var a, b, c int
	triplets := make([]Triangle, 0)

	for a = 1; a < n/3; a++ {
		for b = a + 1; b < n-2*a; b++ {
			c = n - a - b
			if c*c == a*a+b*b {
				triplets = append(triplets, Triangle{a: a, b: b, c: c})
			}
		}
	}

	return triplets
}
