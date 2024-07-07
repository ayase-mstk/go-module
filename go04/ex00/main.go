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
	printStr(os.Args[0])
	ft.PrintRune('\n')
}
