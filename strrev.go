package piscine

func StrRev(s string) string {
	symb := []rune(s)
	for i, j := 0, len(symb)-1; i < j; i, j = i+1, j-1 {
		symb[i], symb[j] = symb[j], symb[i]
		return string(symb)
	}
}
