package piscine

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func numOfSplit(s string) int {
	cnt := 0
	in_word := false
	for _, r := range s {
		if isSpace(r) {
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

func SplitWhiteSpaces(s string) []string {
	var ret []string = make([]string, 0, numOfSplit(s))
	var start int = -1 // wordの中ではなかったら-1にする変数
	for i, r := range s {
		if isSpace(r) {
			if start != -1 {
				ret = append(ret, s[start:i])
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
	}
	if start != -1 {
		ret = append(ret, s[start:])
	}
	return ret
}
