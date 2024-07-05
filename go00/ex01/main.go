package main

import (
	"ft"
)

func main() {
	for c := 'z'; c >= 'a'; c-- {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
