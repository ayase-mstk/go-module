package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrime(-1))
	fmt.Println(piscine.IsPrime(0))
	fmt.Println(piscine.IsPrime(1))
	fmt.Println(piscine.IsPrime(2))
	fmt.Println(piscine.IsPrime(3))
	fmt.Println(piscine.IsPrime(4))
	fmt.Println(piscine.IsPrime(5))
	fmt.Println(piscine.IsPrime(7))
	fmt.Println(piscine.IsPrime(47))
	fmt.Println(piscine.IsPrime(53))
}
