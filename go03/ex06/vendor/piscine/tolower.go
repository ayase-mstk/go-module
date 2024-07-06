package piscine

func sliceLen(sl []rune) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func isUpperCase(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func ToLower(s string) string {
	var sl []rune = []rune(s)
	for i := 0; i < sliceLen(sl); i++ {
		if isUpperCase(sl[i]) {
			sl[i] = sl[i] + ('a' - 'A')
		}
	}
	return string(sl)
}
