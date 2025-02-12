package piscine

func MakeRange(min, max int) []int {
	var ret []int
	if min >= max {
		return ret
	}
	ret = make([]int, max-min)
	for i, n := 0, min; n < max; i, n = i+1, n+1 {
		ret[i] = n
	}
	return ret
}
