package piscine

func RecursiveFactorial(nb int) int {
	if nb <= 1 {
		return 1
	}
	return nb * RecursiveFactorial(n-1)
}
