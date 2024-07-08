package main

import (
	//"math"
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FindNextPrime(-1))
	fmt.Println(piscine.FindNextPrime(4))
	fmt.Println(piscine.FindNextPrime(5))
	fmt.Println(piscine.FindNextPrime(48))
	//array := []int{-1, 0, 1, 2, 3, 4, 5, math.MaxInt32 - 1, math.MaxInt32, math.MaxInt - 24, math.MaxInt}
	//for _, i := range array {
	//	fmt.Printf("[%20d]\t%d\n", i, piscine.FindNextPrime(i))
	//}
}
