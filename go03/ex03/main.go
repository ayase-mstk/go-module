package main

import (
	"fmt"
	"piscine"
	//"strings"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	//fmt.Println(strings.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	//fmt.Println(strings.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
	//fmt.Println(strings.Index("Ola!", "hOl"))
	fmt.Println(piscine.Index("おはよう", "よ"))
	//fmt.Println(strings.Index("おはよう", "よ"))
}
