package piscine

import "ft"

const MIN_INT = ^(int(^uint(0) >> 1))

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

func recursivePrintNbrBase(nbr, base_num int, base []rune) {
	if nbr == 0 {
		return
	}
	recursivePrintNbrBase(nbr/base_num, base_num, base)
	n := nbr % base_num
	ft.PrintRune(base[n])
}

func PrintNbrBase(nbr int, base string) {
	sl_base := []rune(base)
	base_num := sliceLen(sl_base)
	isMinInt := false
	var lastRune rune
	if !isValidBase(sl_base) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	if nbr == 0 {
		ft.PrintRune(sl_base[0])
	}
	if nbr < 0 {
		ft.PrintRune('-')
	}
	if nbr == MIN_INT {
		isMinInt = true
		n := nbr % base_num
		n = -n
		lastRune = sl_base[n]
		nbr /= base_num
	}
	if nbr < 0 {
		nbr = -nbr
	}
	recursivePrintNbrBase(nbr, base_num, sl_base)
	if isMinInt {
		ft.PrintRune(lastRune)
	}
}
