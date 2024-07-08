package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Sqrt(-1))
	fmt.Println(piscine.Sqrt(0))
	fmt.Println(piscine.Sqrt(1))
	fmt.Println(piscine.Sqrt(2))
	fmt.Println(piscine.Sqrt(3))
	fmt.Println(piscine.Sqrt(4))
	fmt.Println(piscine.Sqrt(5))
	fmt.Println(piscine.Sqrt(42 * 42))
	fmt.Println(piscine.Sqrt(3037000499.0 * 3037000499.0))
	//fmt.Println(piscine.Sqrt(3037000500 * 3037000499))
}
