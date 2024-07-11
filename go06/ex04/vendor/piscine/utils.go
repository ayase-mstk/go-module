package piscine

func sliceLen[T any](sl []T) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func strLen(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func isNumeric(r rune) bool {
	return '0' <= r && r <= '9'
}

func isNumericWord(s string) bool {
	sl := []rune(s)
	for i := 0; i < sliceLen(sl); i++ {
		if !isNumeric(sl[i]) {
			return false
		}
	}
	return true
}
