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
	args := os.Args[1:]
	l := sliceLen(args)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for i := 0; i < l; i++ {
		printStr(args[i])
		ft.PrintRune('\n')
	}
}
