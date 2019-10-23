/*package main

import "fmt"*/

package piscine

func ToLower(s string) string {

	str := []rune(s)
	for i, j := range str {
		if j >= 'A' && j <= 'Z' {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}

/*func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}*/
