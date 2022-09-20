package problems

// AllDivisorsOfNumber n returns slice of all its divisors
func AllDivisorsOfNumber(n int) []int {
	divisors := []int{1}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	divisors = append(divisors, n)

	return divisors
}
