package piscine

import (
	"ft"
)

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
	if nbr < 0 {
		ft.PrintRune('-')
		nbr = -nbr
	}
	if nbr == 0 {
		return
	}
	recursivePrintNbrBase(nbr/base_num, base_num, base)
	n := nbr % base_num
	ft.PrintRune(base[n])
}

func PrintNbrBase(nbr int, base string) {
	sl_base := []rune(base)
	if !isValidBase(sl_base) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	base_num := sliceLen(sl_base)
	recursivePrintNbrBase(nbr, base_num, sl_base)
}
