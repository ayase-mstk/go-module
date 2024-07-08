package piscine

func Sqrt(nb int) int {
	ret := 0
	for i := 1; i <= nb/i; i++ {
		if i == nb/i && nb%i == 0 {
			ret = i
			break
		}
	}
	return ret
}
