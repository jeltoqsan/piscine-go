package piscine

import "fmt"

func PrintStr(str string) {
	for _, x := range str {
		fmt.Print(string(x))
	}
}
