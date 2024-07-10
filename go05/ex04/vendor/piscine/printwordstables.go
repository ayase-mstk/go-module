package piscine

import "ft"

func printWord(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func PrintWordsTables(a []string) {
	for _, s := range a {
		printWord(s)
		ft.PrintRune('\n')
	}
}
