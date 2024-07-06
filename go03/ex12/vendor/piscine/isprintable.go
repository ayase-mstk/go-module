package piscine

func sliceLen(sl []rune) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func isPrintable(r rune) bool {
	return 31 < r && r < 127
}

func IsPrintable(s string) bool {
	sl := []rune(s)
	for i := 0; i < sliceLen(sl); i++ {
		if !isPrintable(sl[i]) {
			return false
		}
	}
	return true
}
