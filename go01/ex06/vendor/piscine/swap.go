package piscine

func Swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
