package piscine

func sliceLen(s []rune) int {
	var len int = 0
	for range s {
		len++
	}
	return len
}

func StrRev(s string) string {
	// 文字列はbytesごとのアクセス
	// runeスライスに変換することでrune単位で操作できる
	ret := []rune(s)

	for i, j := 0, sliceLen(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}
