package piscine

import (
	"ft"
)

func PrintStr(str string) {
	for _, r := range str {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
