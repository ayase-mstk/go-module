package main

import (
	"ft"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		printStr(arg)
		ft.PrintRune('\n')
	}
}
