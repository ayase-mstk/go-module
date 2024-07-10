package piscine

import "ft"

const (
	MinInt          = -1 << (64 - 1)
	InvalidBaseChar = "NV"
)

func sliceLen(sl []rune) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func isValidBase(base []rune) bool {
	if sliceLen(base) < 2 {
		return false
	}
	charMap := make(map[rune]bool)
	for _, c := range base {
		if c == '+' || c == '-' || charMap[c] {
			return false
		}
		charMap[c] = true
	}
	return true
}

func PrintNbrBaseRecursive(nbr, baseNum int, base []rune) {
	if nbr == 0 {
		return
	}
	PrintNbrBaseRecursive(nbr/baseNum, baseNum, base)
	ft.PrintRune(base[nbr%baseNum])
}

func HandleMinInt(nbr int, baseNum int, base []rune) (int, rune) {
	isMinInt := nbr == MinInt
	var lastRune rune
	if isMinInt {
		n := -(nbr % baseNum)
		lastRune = base[n]
		nbr /= baseNum
	}
	return nbr, lastRune
}

func PrintNbrBase(nbr int, base string) {
	slBase := []rune(base)
	baseNum := sliceLen(slBase)

	if !isValidBase(slBase) {
		for _, r := range InvalidBaseChar {
			ft.PrintRune(r)
		}
		return
	}

	if nbr == 0 {
		ft.PrintRune(slBase[0])
		return
	}

	if nbr < 0 {
		ft.PrintRune('-')
	}

	nbr, lastRune := HandleMinInt(nbr, baseNum, slBase)

	if nbr < 0 {
		nbr = -nbr
	}

	PrintNbrBaseRecursive(nbr, baseNum, slBase)

	if lastRune != 0 {
		ft.PrintRune(lastRune)
	}
}
