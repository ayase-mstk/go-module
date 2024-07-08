package piscine

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	for i := 5; i*i <= nb; i += 6 { // i = 6n - 1でloop
		if nb%i == 0 || nb%(i+2) == 0 { // 6n-1と6n+1で計算する
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
