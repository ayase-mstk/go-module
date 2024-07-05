package piscine

import (
	"ft"
)

func ftLen(s string) int {
	var len int = 0
	for range s {
		len++
	}
	return len
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func generateCombN(n, start int, arr string) {
	// 終了条件
	// prefix == nのものしかresultsに追加されない
	if ftLen(arr) == n {
		printStr(arr)
		if rune(arr[0]) != rune(10-n+'0') {
			printStr(", ")
		} else {
			ft.PrintRune('\n')
		}
		return
	}

	for i := start; i < 10; i++ {
		generateCombN(n, i+1, arr+string(i+'0'))
	}
}

func PrintCombN(n int) {
	if n <= 0 || 10 <= n {
		printStr("Error: n should be between 1 and 9\n")
		return
	}

	generateCombN(n, 0, "")
}
