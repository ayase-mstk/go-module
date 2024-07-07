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
	args := os.Args
	l := sliceLen(args)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for i := 0; i < l; i++ {
		printStr(os.Args[i])
		ft.PrintRune('\n')
	}
}
