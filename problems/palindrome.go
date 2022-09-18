package problems

// IsPalindrome function check if, the number reversed is same as original number or not
// eg:
// 323 reverse is 323 = true
// 9 reverse is 9 = true
// 32 reverse is 23 = false
func IsPalindrome(n int) bool {
	rev := 0
	og := n // persist the original value
	for n > 0 {
		rev = (rev * 10) + (n % 10)
		n = n / 10
	}
	return rev == og
}
