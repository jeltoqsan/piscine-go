package piscine

func IterativeFactorial(nb int) int { // принимает инт, отдает инт
	if nb >= 0 && nb <= 20 { // 20 - потому что дальше выходит за пределы инт
		result := 1
		for i := 1; i <= nb; i++ { // делаться ай пока меньше или равно нб
			result = result * i // result *= i
		}
		return result
	} else {
		return 0
	}
}
