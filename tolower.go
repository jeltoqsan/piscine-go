package piscine

func Tolower(s string) string {

	str := []rune(s)
	for i, j := range str {
		if j >= 'Z' && j <= 'A' {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}
