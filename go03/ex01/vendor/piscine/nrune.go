package piscine

func sliceLen(sl []rune) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func NRune(s string, n int) rune {
	var ret []rune = []rune(s)

	if n <= 0 || sliceLen(ret) < n {
		return 0
	}

	return ret[n-1]
}
