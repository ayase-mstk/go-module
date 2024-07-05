package piscine

func UltimateDivMod(a, b *int) {
	tmp := *a / *b
	*b = *a % *b
	*a = tmp
}
