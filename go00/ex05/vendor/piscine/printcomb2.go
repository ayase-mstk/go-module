package piscine

import (
	"ft"
)

func PrintNum(num int) {
	c := num / 10
	ft.PrintRune(rune(c + '0'))
	c = num % 10
	ft.PrintRune(rune(c + '0'))
}

func PrintComb2() {
	for i := 0; i < 99; i++ {
		for j := i + 1; j < 100; j++ {
			PrintNum(i)
			ft.PrintRune(' ')
			PrintNum(j)
			if !(i == 98) {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			} else {
				ft.PrintRune('\n')
			}
		}
	}
}
