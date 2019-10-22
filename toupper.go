/*package main

import "fmt"*/

package piscine

func ToUpper(s string) string {

	str := []rune(s)
	for i, j := range str {
		if j > 'a' && j < 'z' {
			str[i] = str[i] - 32
		}
	}
	return string(str)
}

/*func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}*/
