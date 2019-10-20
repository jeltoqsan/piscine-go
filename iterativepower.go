package piscine

func IterativePower(nb int, power int) int {
	qwerty := 1
	if power <= 0 {
		return 0
	} else {
		for i := 0; i < power; i++ {
			qwerty *= nb
		}
		return qwerty
	}
}
