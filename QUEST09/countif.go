package piscine

func CountIf(f func(string) bool, tab []string) int {
	x := 0
	for _, char := range tab {
		if f(char) {
			x = x + 1
		}
	}
	return x
}
