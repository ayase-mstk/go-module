package piscine

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func atoi(s string) int {
	if strLen(s) == 0 {
		return 0
	}

	const MAX_INT = int(^uint(0) >> 1)
	const MIN_INT = ^(int(^uint(0) >> 1))
	ret := 0
	sign := 1

	if rune(s[0]) == '-' || rune(s[0]) == '+' {
		if rune(s[0]) == '-' {
			sign = -1
		}
		s = s[1:]
	}

	for _, r := range s {
		if !isDigit(r) {
			return 0
		}
		if sign == 1 && (ret > MAX_INT/10 || (ret == MAX_INT/10 && int(r-'0') > MAX_INT%10)) {
			return MAX_INT
		}
		if sign == -1 && (ret > MAX_INT/10 || (ret == MAX_INT/10 && int(r-'0') > MAX_INT%10)) {
			return MIN_INT
		}
		ret *= 10
		ret += int(r - '0')
	}
	return ret * sign
}
