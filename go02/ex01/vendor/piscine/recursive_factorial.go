package piscine

const INT_MAX = int(^uint(0) >> 1)

func RecursiveFactorial(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb <= 0 {
		return 0
	}

	ret := RecursiveFactorial(nb - 1)
	if INT_MAX/nb < ret {
		return 0
	}

	return nb * ret
}
