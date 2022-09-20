package problems

// PrimeFactors return prime factors of given number n
func PrimeFactors(n int) []int {
	if n <= 1 {
		return []int{}
	}
	facts := make([]int, 0)
	for i := 2; i*i <= n; i++ {

		for n%i == 0 {
			facts = append(facts, i)
			n = n / i
		}
	}

	if n > 1 {
		facts = append(facts, n)
	}
	return facts
}
