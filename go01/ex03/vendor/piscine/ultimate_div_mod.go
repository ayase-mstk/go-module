package piscine

func UltimateDivMod(a, b *int) {
	*a, *b = *a / *b, *a%*b
}
