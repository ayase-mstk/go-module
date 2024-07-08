package piscine

//import "fmt"

func Split(s, sep string) []string {
	var ret []string = make([]string, 0, countSplit(s, sep))
	var sep_len int = strLen(sep)
	var start int = -1

	var rune_s []rune = []rune(s)

	for i := 0; i < strLen(s); {
		//fmt.Printf("%#U\n", rune_s[i])
		if string(rune_s[i:i+sep_len]) == sep {
			if start != -1 {
				ret = append(ret, string(rune_s[start:i]))
				start = -1
			} else {
				ret = append(ret, string(""))
			}
			i += sep_len
		} else {
			if start == -1 {
				start = i
			}
			i++
		}
	}
	if start != -1 {
		ret = append(ret, string(rune_s[start:]))
	} else {
		ret = append(ret, string(""))
	}
	return ret
}

func countSplit(s, sep string) int {
	var cnt int = 0
	var sep_len int = strLen(sep)
	var rune_s []rune = []rune(s)
	var in_word bool = false

	for i := 0; i < strLen(s)-sep_len; i++ {
		if string(rune_s[i:i+sep_len]) == sep {
			if in_word {
				cnt++
				in_word = false
			}
		} else {
			in_word = true
		}
	}
	if in_word {
		cnt++
	}
	return cnt
}

func strLen(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}
