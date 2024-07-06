package piscine

func sliceLen(sl []rune) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func isNumeric(r rune) bool {
	return '0' <= r && r <= '9'
}

func isLowerCase(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isUpperCase(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func isAlphaNumeric(r rune) bool {
	return isLowerCase(r) || isUpperCase(r) || isNumeric(r)
}

func IsAlpha(s string) bool {
	sl := []rune(s)
	for i := 0; i < sliceLen(sl); i++ {
		if !isAlphaNumeric(sl[i]) {
			return false
		}
	}
	return true
}
