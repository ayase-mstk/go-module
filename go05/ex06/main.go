package main

import (
	"fmt"
	"piscine"
	"strings"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = ",a,b,c"
	fmt.Printf("%#v\n", piscine.Split(s, ","))
	fmt.Printf("%#v\n", strings.Split(s, ","))
	s = ""
	fmt.Printf("%#v\n", piscine.Split(s, ","))
	fmt.Printf("%#v\n", strings.Split(s, ","))
	s = ",,,,,,,"
	fmt.Printf("%#v\n", piscine.Split(s, ","))
	fmt.Printf("%#v\n", strings.Split(s, ","))
}
