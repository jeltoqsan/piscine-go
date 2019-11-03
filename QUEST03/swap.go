package piscine

func Swap(a, b *int) {
	var x int
	x = *a
	*a = *b
	*b = x
}
