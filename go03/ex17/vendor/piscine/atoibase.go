package piscine

func sliceLen(sl []rune) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func isValidBase(base []rune) bool {
	// at least 2 characters
	// each character must be unique
	// should not contain sign
	if sliceLen(base) < 2 {
		return false
	}
	charMap := map[rune]bool{}
	for _, c := range base {
		if c == '+' || c == '-' || charMap[c] {
			return false
		}
		charMap[c] = true
	}
	return true
}

func isInBase(s string, base []rune) bool {
	for _, c := range s {
		is_in_base := false
		for _, base_char := range base {
			if c == base_char {
				is_in_base = true
			}
		}
		if !is_in_base {
			return false
		}
	}
	return true
}

func nbrBase(nbr rune, base []rune) int {
	for i, r := range base {
		if nbr == r {
			return i
		}
	}
	return 0
}

func AtoiBase(s string, base string) int {
	sl_base := []rune(base)
	if !isValidBase(sl_base) || !isInBase(s, sl_base) {
		return 0
	}

	ret := 0
	base_num := sliceLen(sl_base)
	for _, c := range s {
		ret *= base_num
		ret += nbrBase(c, sl_base)
	}
	return ret
}
