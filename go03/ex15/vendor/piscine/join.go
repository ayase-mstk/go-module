package piscine

func sliceLen(sl []string) int {
	l := 0
	for range sl {
		l++
	}
	return l
}

func Join(strs []string, sep string) string {
	var ret string
	for i := 0; i < sliceLen(strs); i++ {
		if i != 0 {
			ret = ret + sep
		}
		ret = ret + strs[i]
	}
	return ret
}
