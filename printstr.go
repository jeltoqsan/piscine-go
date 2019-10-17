package piscine

func PrintRune(str string) {
	for _, x := range str {
		fmt.PrintRune(string(x))
	}
}
