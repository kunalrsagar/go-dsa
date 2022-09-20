package problems

// PrimeFactors return prime factors of given number n
func PrimeFactors(n int) []int {
	if n <= 1 {
		return []int{}
	}
	facts := make([]int, 0)

	for n%2 == 0 {
		facts = append(facts, 2)
		n = n / 2
	}

	for n%3 == 0 {
		facts = append(facts, 3)
		n = n / 3
	}

	for i := 5; i*i <= n; i += 6 {

		for n%i == 0 {
			facts = append(facts, i)
			n = n / i
		}

		for n%(i+2) == 0 {
			facts = append(facts, i+2)
			n = n / (i + 2)
		}
	}

	if n > 1 {
		facts = append(facts, n)
	}
	return facts
}
