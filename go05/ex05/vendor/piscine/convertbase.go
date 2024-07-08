package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return nbrBase(atoiBase(nbr, baseFrom), baseTo)
}

func nbrBase(decimal int, baseTo string) string {
	var base_num int = strLen(baseTo)
	var ret []rune = make([]rune, 0, countDigit(decimal, base_num))
	var base []rune = []rune(baseTo)

	if decimal == 0 {
		return string(base[0])
	}

	for decimal > 0 {
		ret = append(ret, base[decimal%base_num])
		decimal /= base_num
	}
	for i, j := 0, sliceLen(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}

func countDigit(nbr, digit int) int {
	ret := 0
	for nbr > 0 {
		ret++
		nbr /= digit
	}
	return ret
}

func atoiBase(s string, base string) int {
	sl_base := []rune(base)
	//if !isValidBase(sl_base) || !isInBase(s, sl_base) {
	//	return 0
	//}

	ret := 0
	base_num := sliceLen(sl_base)
	for _, c := range s {
		ret *= base_num
		ret += indexBase(c, sl_base)
	}
	return ret
}

func strLen(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

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

func indexBase(nbr rune, base []rune) int {
	for i, r := range base {
		if nbr == r {
			return i
		}
	}
	return 0
}
