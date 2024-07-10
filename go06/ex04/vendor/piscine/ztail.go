package piscine

import "os"

func argsValidation() (int, error) {
	argc := stringSliceLen(os.Args[1:])
	args := os.Args[1:]
	var err CustomError

	if argc == 1 && args[0] != "-c" {
		err.setMsg("ztail: only the -c option is supported.")
		return 0, err
	}

	if argc == 1 {
		err.setMsg("ztail: option requires an argument -- 'c'")
		return 0, err
	}

	if !isNumericWord(args[1]) {
		err.setMsg("ztail: invalid number of bytes: '" + args[1] + "'")
		return 0, err
	}

	return atoi(args[1]), nil
}

func ztailStdIn() {
	var buffer [256]byte
	for {
		n, err := os.Stdin.Read(buffer[:])
		if n == 0 { // finish Read or EOF
			break
		}
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			continue
		}
		os.Stdout.Write(buffer[:n])
	}
}

func ztailReadFile(fileName string, tail int) bool {
	data, err := os.ReadFile(fileName)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return false
	}
	os.Stdout.Write(data[byteSliceLen(data)-tail:])
	return true
}

func printHead(fileName string) {
  if strLen(fileName) == 1 && fileName == "-" {
    fileName = "standard input"
  }
	output := "==> " + fileName + " <==\n"
	os.Stdout.WriteString(output)
}

func Ztail() {
	tail, err := argsValidation()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	args := os.Args[3:]
	argc := stringSliceLen(os.Args[3:])
	exitCode := 0
	var n int = 0
	switch argc {
	case 0:
		ztailStdIn()
		break

	case 1:
		if strLen(args[n]) == 1 && rune(args[n][0]) == '-' {
			ztailStdIn()
		} else {
			if !ztailReadFile(args[n], tail) {
				exitCode = 1
			}
		}
		break

	default:
		for n < argc {
			printHead(args[n])
			if strLen(args[n]) == 1 && rune(args[n][0]) == '-' {
				ztailStdIn()
			} else {
				if !ztailReadFile(args[n], tail) {
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
