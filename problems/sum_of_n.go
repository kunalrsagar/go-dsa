// Package problems have problems and solutions
// Sum of n natural number problem for Analysis of Algorithms.
package problems

// SumOfN function is used to sum n natural numbers
func SumOfN(n int) int {
	var sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// SumOfNV2 function uses math formula to sum n natural numbers
func SumOfNV2(n int) int {
	return n * (n + 1) / 2
}
