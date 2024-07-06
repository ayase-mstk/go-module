package piscine

func NRune(s string, n int) rune {
	var ret []rune = []rune(s)

	if n <= 0 || len(ret) < n {
		return 0
	}

	return ret[n-1]
}
