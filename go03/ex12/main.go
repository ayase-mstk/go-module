package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("Hello"))
	fmt.Println(piscine.IsPrintable("Hello\n"))
	fmt.Println(piscine.IsPrintable(" "))
	fmt.Println(piscine.IsPrintable(""))
}
