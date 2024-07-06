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

func ToUpper(s string) string {
	var sl []rune = []rune(s)
	for i := 0; i < sliceLen(sl); i++ {
		if isLowerCase(sl[i]) {
			sl[i] = sl[i] - ('a' - 'A')
		}
	}
	return string(sl)
}
