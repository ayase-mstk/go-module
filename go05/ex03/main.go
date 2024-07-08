package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" Hello how are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you? "))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("\t Hello\nhow\tare you? \n \t"))
}
