package piscine

func DivMod(a, b int, div, mod *int) {
  *div, *mod = a / b, a % b
}
