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

func IsNumeric(s string) bool {
	sl := []rune(s)
	for i := 0; i < sliceLen(sl); i++ {
		if !isNumeric(sl[i]) {
			return false
		}
	}
	return true
}
