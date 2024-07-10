package main

import (
	"fmt"
	"piscine"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
	test = []string{"Hello", "", "", "you?"}
	fmt.Println(piscine.ConcatParams(test))
}
