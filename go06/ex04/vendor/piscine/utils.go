package piscine

func stringSliceLen(s []string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func byteSliceLen(s []byte) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func runeSliceLen(s []rune) int {
	l := 0
	for range s {
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
	for i := 0; i < runeSliceLen(sl); i++ {
		if !isNumeric(sl[i]) {
			return false
		}
	}
	return true
}
