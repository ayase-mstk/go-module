package main

import (
	"ft"
	"math"
	"piscine"
)

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(0, "choumi")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(0, "あいう")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(0, "")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(math.MinInt, "0123456789")
	ft.PrintRune('\n')
}
