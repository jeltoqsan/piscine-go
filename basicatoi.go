package piscine

func BasicAtoi(s string) int {
	var x int

	for _, y := range s {
		x *= 10
		if y == 48 {
			x += 0
		} else if y == 49 {
			x += 1
		} else if y == 50 {
			x += 2
		} else if y == 51 {
			x += 3
		} else if y == 52 {
			x += 4
		} else if y == 53 {
			x += 5
		} else if y == 54 {
			x += 6
		} else if y == 55 {
			x += 7
		} else if y == 56 {
			x += 8
		} else if y == 57 {
			x += 9
		}
	}
	return x
}
