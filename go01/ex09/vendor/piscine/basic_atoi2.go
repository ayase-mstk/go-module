package piscine

func strLen(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func BasicAtoi2(s string) int {
	if strLen(s) == 0 {
		return 0
	}
	const INT_MAX = int(^uint(0) >> 1)
	ret := 0

	for _, r := range s {
		if !isDigit(r) {
			return 0
		}
		if ret > INT_MAX/10 || (ret == INT_MAX/10 && int(r-'0') > INT_MAX%10) {
			return INT_MAX
		}
		ret *= 10
		ret += int(r - '0')
	}
	return ret
}
