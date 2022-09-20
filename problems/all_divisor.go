package problems

import "sort"

// AllDivisorsOfNumber n returns slice of all its divisors
func AllDivisorsOfNumber(n int) []int {
	divisors := []int{1}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != (n / i) {
				divisors = append(divisors, n/i)
			}
		}
	}

	divisors = append(divisors, n)
	sort.Ints(divisors)
	return divisors
}
