package problems

// Gcd - Greatest Common divisor
// using Euclidean Algorithm
// https://www.geeksforgeeks.org/euclidean-algorithms-basic-and-extended/
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
