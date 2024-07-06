package piscine

func LastRune(s string) rune {
	var ret []rune = []rune(s)
	if len(ret) == 0 {
		return 0
	}
	return ret[len(s)-1]
}
