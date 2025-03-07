package leetcode

func closestPrimes(left int, right int) []int {

	var ans = []int{-1, -1}
	var primes []int

	for i := left; i <= right; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	lPrimes := len(primes)
	if lPrimes < 2 {
		return ans
	}

	minDif := primes[1] - primes[0]
	ans = []int{primes[0], primes[1]}

	for i := 1; i < lPrimes-1; i++ {
		if minDif > primes[i+1]-primes[i] {
			minDif = primes[i+1] - primes[i]
			ans = []int{primes[i], primes[i+1]}
		}
	}

	return ans
}

func isPrime(num int) bool {

	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func ClosestPrimes() {
	closestPrimes(10, 19)
}
