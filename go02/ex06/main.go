package main

import (
	//"math"
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrime(-1))
	fmt.Println(piscine.IsPrime(0))
	fmt.Println(piscine.IsPrime(2))
	fmt.Println(piscine.IsPrime(3))
	fmt.Println(piscine.IsPrime(5))
	fmt.Println(piscine.IsPrime(7))
	fmt.Println(piscine.IsPrime(53))
	//array := []int{-1, 0, 1, 2, 3, 4, 5, math.MaxInt32 - 1, math.MaxInt32, math.MaxInt - 24, math.MaxInt}
	//for _, i := range array {
	//	fmt.Printf("[%20d]\t%d\n", i, piscine.IsPrime(i))
	//}
}
