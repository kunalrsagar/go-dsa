package problems

// TrailingZeroFact is a function that returns
// the count of zero tailing the factorial of given number
// Eg: 5! = 120 = 1 trailing zero
func TrailingZeroFact(n int) int {
	zero := 0
	for i := 5; i <= n; i = i * 5 {
		zero += n / i
	}

	return zero
}
