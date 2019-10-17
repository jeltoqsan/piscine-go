package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a
	d := *b
	*a = (d / c)
	*b = (d % c)
}
