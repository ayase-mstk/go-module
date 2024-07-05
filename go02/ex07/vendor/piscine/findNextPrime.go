package piscine

func sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}

	left, right := 1, nb
	var ans int
	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= nb {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func isPrime(nb int) bool {
	if nb <= 2 {
		return false
	}
	for i := 2; i <= sqrt(nb); i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for !isPrime(nb) {
		nb++
	}
	return nb
}
