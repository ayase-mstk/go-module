package piscine

import (
	"ft"
)

func PrintComb() {
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			for k := j + 1; k < 10; k++ {
				ft.PrintRune(rune(i + '0'))
				ft.PrintRune(rune(j + '0'))
				ft.PrintRune(rune(k + '0'))
				if !(i == 7) {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				} else {
					ft.PrintRune('\n')
				}
			}
		}
	}
}
