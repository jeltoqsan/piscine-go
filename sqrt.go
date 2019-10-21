package piscine

func Sqrt(nb int) int {

	for i := 1; i <= nb; i++ {
		if i*i > nb {
			return 0
		} else if i*i == nb {
			return i
		}
	}
	return 0
}
