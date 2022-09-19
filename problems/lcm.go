package problems

// LCM function returns Largest common multiple of given 2 numbers a and b
func LCM(a, b int) int {
	return (a * b) / Gcd(a, b)
}
