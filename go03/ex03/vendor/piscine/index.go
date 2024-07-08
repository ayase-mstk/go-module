package piscine

func strLen(s string) int {
	bytes := []byte(s)
	l := 0
	for range bytes {
		l++
	}
	return l
}

func Index(s string, toFind string) int {
	n := strLen(toFind)

	switch {
	case n == 0:
		return 0
	case n == 1:
		for i := 0; i < strLen(s); i++ {
			if s[i] == toFind[0] {
				return i
			}
		}
		return -1
	case n == strLen(s):
		if s == toFind {
			return 0
		}
		return -1
	case n > strLen(s):
		return -1
	default:
		for i := 0; i < strLen(s)-n; i++ {
			if s[i:i+n] == toFind {
				return i
			}
		}
		return -1
	}
}
