package piscine

const INT_MAX = int(^uint(0) >> 1)

func IterativeFactorial(nb int) int {
	if nb <= 0 {
		return 0
	}

	var ret int = 1
	for i := 1; i <= nb; i++ {
		if INT_MAX/i < ret {
			return 0
		}
		ret *= i
	}
	return ret
}
