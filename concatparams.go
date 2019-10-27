package piscine

/*package main

import "fmt"*/

func ConcatParams(args []string) string {
	str := ""
	for i, r := range args {
		if i == 0 {
			str += r
		} else {
			str += "\n" + r
		}
	}
	return str
}

/*func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}*/
