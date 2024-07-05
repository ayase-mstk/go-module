package piscine

func BasicAtoi(s string) int {
	ret := 0

	for _, r := range s {
		ret *= 10
		ret += int(r - '0')
	}
	return ret
}
