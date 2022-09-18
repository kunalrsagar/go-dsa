package problems

// CountDigits return the no. of digit n has
// eg: 5420 returns 4,
// eg: 32 returns 2
func CountDigits(n int) int {
	if n == 0 {
		return 1
	}
	c := 0

	for ; n > 0; n = n / 10 {
		c++
	}

	return c
}
