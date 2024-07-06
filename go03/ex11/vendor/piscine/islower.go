package piscine

func sliceLen(sl []rune) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func isLowerCase(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func IsLower(s string) bool {
	sl := []rune(s)
	for i := 0; i < sliceLen(sl); i++ {
		if !isLowerCase(sl[i]) {
			return false
		}
	}
	return true
}
