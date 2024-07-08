package main

import (
	"ft"
	"os"
)

func findLastSlash(s string) int {
	n := 0
	for i, r := range s {
		if r == '/' {
			n = i
		}
	}
	return n
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func main() {
	n := findLastSlash(os.Args[0])
	printStr(os.Args[0][n+1:])
	ft.PrintRune('\n')
}
