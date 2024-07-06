package piscine

func FirstRune(s string) rune {
	sl := []rune(s)
	if len(s) == 0 {
		return 0
	}
	return sl[0]
}
