package piscine

func sliceLen(sl []string) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func BasicJoin(elems []string) string {
	var ret string
	for i := 0; i < sliceLen(elems); i++ {
		ret = ret + elems[i]
	}
	return ret
}
