package piscine

const MaxPrime = int(^uint(0)>>1) - 24

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	for i := 2; i <= nb/i; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb > MaxPrime {
		return -1
	}
	for !isPrime(nb) {
		nb++
	}
	return nb
}
