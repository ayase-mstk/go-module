package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.IterativeFactorial(arg))
	arg = -1
	fmt.Println(piscine.IterativeFactorial(arg))
	arg = 0
	fmt.Println(piscine.IterativeFactorial(arg))
	arg = 20
	fmt.Println(piscine.IterativeFactorial(arg))
	arg = 21
	fmt.Println(piscine.IterativeFactorial(arg))
}
