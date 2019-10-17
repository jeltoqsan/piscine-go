package piscine

func DivMod(a, b int, div, mod *int) {

	с := (a / b)
	d := a % b
	*div = c
	*mod = d
}
