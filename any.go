package piscine

/*import (
	"fmt"

	piscine ".."
)*/

func Any(f func(string) bool, arr []string) bool {
	for _, char := range arr {
		if f(char) {
			return true
		}
	}
	return false
}

/*func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "is", "4", "you"}

	result1 := piscine.Any(piscine.IsNumeric, tab1)
	result2 := piscine.Any(piscine.IsNumeric, tab2)

	fmt.Println(result1)
	fmt.Println(result2)
}*/
