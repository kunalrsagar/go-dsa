package problems

// SieveOfEratosthenes function returns slice of all prime number less than input n
func SieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if IsPrime(i) {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	primeNumbers := make([]int, 0)

	for i, v := range primes {
		if v {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}
