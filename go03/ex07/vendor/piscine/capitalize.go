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

func isUpperCase(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func isNumeric(r rune) bool {
	return '0' <= r && r <= '9'
}

func isAlphaNumeric(r rune) bool {
	return isLowerCase(r) || isUpperCase(r) || isNumeric(r)
}

func Capitalize(s string) string {
	var sl []rune = []rune(s)
	for i := 0; i < sliceLen(sl); i++ {
		if i == 0 || (i != 0 && !isAlphaNumeric(sl[i-1])) {
			if isLowerCase(sl[i]) {
				sl[i] = sl[i] - ('a' - 'A')
			}
		} else if isUpperCase(sl[i]) {
			sl[i] = sl[i] + ('a' - 'A')
		}
	}
	return string(sl)
}
