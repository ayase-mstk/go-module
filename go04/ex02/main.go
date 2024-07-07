package main

import (
	"ft"
	"os"
)

func sliceLen(s []string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func main() {
	l := sliceLen(os.Args)
	for i := l - 1; i > 0; i-- {
		printStr(os.Args[i])
		ft.PrintRune('\n')
	}
}
