package main

import (
	"fmt"
	"piscine"
	//"strings"
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	//fmt.Println(strings.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Salut!", "lut!"))
	//fmt.Println(strings.Compare("Salut!", "lut!"))
	fmt.Println(piscine.Compare("Ola!", "Ol"))
	//fmt.Println(strings.Compare("Ola!", "Ol"))
	fmt.Println(piscine.Compare("おはよう", "おは"))
	//fmt.Println(strings.Compare("おはよう", "おは"))
}
