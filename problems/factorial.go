package problems

// Factorial returns factorial of a given number
// eg:
// 6, 720
// 5, 120
// 0, 1
func Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}
