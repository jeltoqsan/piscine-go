package piscine

func UltimateDivMod(a *int, b *int) {
	x := *a
	*a = *a / *b
	*b = x % *b
}
