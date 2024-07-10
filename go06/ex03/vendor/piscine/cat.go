package piscine

import "os"

func sliceLen(s []string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func strLen(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func catStdIn() {
	var buffer [256]byte
	for {
		n, err := os.Stdin.Read(buffer[:])
		if n == 0 { // finish Read or EOF
			break
		}
		if err != nil {
			os.Stderr.WriteString("Error : " + err.Error() + "\n")
      continue
		}
		os.Stdout.Write(buffer[:n])
	}
}

func catFile(fileName string) bool {
	data, err := os.ReadFile(fileName)
	if err != nil {
		os.Stderr.WriteString("Error : " + err.Error() + "\n")
		return false
	}
	os.Stdout.Write(data)
	return true
}

func Cat() {
	args := os.Args[1:]
	argc := sliceLen(args)
	exitCode := 0
	switch argc {
	case 0:
		catStdIn()
		break
	default:
		var n int = 0
		for n < argc {
			if strLen(args[n]) == 1 && rune(args[n][0]) == '-' {
				catStdIn()
			} else {
				if !catFile(args[n]) {
					exitCode = 1
				}
			}
			n++
		}
		break
	}
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
