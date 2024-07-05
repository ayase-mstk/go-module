package piscine

func Sqrt(nb int) int {
	ret := 0
	for i := 1; i*i <= nb; i++ {
		if nb%i == 0 {
			ret = i
		}
	}
	return ret
}
