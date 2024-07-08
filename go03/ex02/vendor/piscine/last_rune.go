package piscine

func sliceLen(sl []rune) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func LastRune(s string) rune {
	var ret []rune = []rune(s)
	if sliceLen(ret) == 0 {
		return 0
	}
	return ret[sliceLen(ret)-1]
}
