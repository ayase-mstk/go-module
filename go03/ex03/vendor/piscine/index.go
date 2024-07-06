package piscine

func Index(s string, toFind string) int {
	n := len(toFind)

	switch {
	case n == 0:
		return 0
	case n == 1:
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[0] {
				return i
			}
		}
		return -1
	case n == len(s):
		if s == toFind {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	default:
		for i := 0; i < len(s)-n; i++ {
			if s[i:i+n] == toFind {
				return i
			}
		}
		return -1
	}
}
