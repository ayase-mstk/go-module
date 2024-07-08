package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	ret := 1
	for i := 0; i < power; i++ {
		ret *= nb
	}
	return ret
}
