package main

import "os"

var NoArgsMsg = []byte("File name missing\n")
var TooManyArgsMsg = []byte("Too many arguments\n")

func strLen(s []string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func displayFile(argc int, s string) {
	data, err := os.ReadFile(s)
	if err != nil {
		os.Stderr.WriteString("Error : " + err.Error() + "\n")
		return
	}
	os.Stdout.Write(data)
}

func main() {
	argc := strLen(os.Args[1:])
	if argc == 0 {
		os.Stdout.Write(NoArgsMsg[:])
		return
	}
	if argc > 1 {
		os.Stdout.Write(TooManyArgsMsg[:])
		return
	}
	displayFile(argc, os.Args[1])
}
