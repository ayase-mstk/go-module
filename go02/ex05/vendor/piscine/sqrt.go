package piscine

func Sqrt(nb int) int {
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return i
		}
	}
	return 0
}
