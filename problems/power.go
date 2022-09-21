package problems

// Power function return power of x raise to n
func Power(x, n int) int {

	if n == 0 {
		return 1
	}

	tmp := Power(x, n/2)
	tmp = tmp * tmp

	if n%2 == 0 {
		return tmp
	}

	return tmp * x
}
