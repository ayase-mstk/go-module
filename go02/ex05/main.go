package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Sqrt(4))
	fmt.Println(piscine.Sqrt(3))
	fmt.Println(piscine.Sqrt(15 * 15))
	fmt.Println(piscine.Sqrt(42 * 42))
	fmt.Println(piscine.Sqrt(0))
	fmt.Println(piscine.Sqrt(-9))
}
