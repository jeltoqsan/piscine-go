package piscine

imp "github.com/01-edu/z01"

func PrintRune(str string) {
	for _, x := range str {
		fmt.PrintRune(string(x))
	}
}
