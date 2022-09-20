package problems

// AllDivisorsOfNumber n returns slice of all its divisors
func AllDivisorsOfNumber(n int) []int {
	divisors := []int{}
	i := 1
	for ; i*i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	for ; i >= 1; i-- {
		if n%i == 0 {
			divisors = append(divisors, n/i)
		}
	}

	//divisors = append(divisors, n)
	return divisors
}
