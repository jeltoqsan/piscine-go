package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return FindNextPrime(nb + 1)
			}
		}
		return nb
	} else if nb == 0 {
		return 2
	} else if nb == 1 {
		return 2
	} else {
		return 2
	}
}
