package piscine

/*package main

import "fmt"*/

func SplitWhiteSpaces(str string) []string {
	counter := 0
	for _, z := range str {
		if z == ' ' || z == '\n' || z == '\t' {
			counter++
		}
	}
	res := make([]string, counter+1)
	j := 0
	last := 0
	for i, z := range str {
		if z == ' ' || z == '\n' || z == '\t' {
			if str[last:i] != "" {
				res[j] = str[last:i]
				j++
			}
			last = i + 1
		}
	}
	res[j] = str[last:]
	c := 0
	for _, z := range res {
		if z != "" {
			c++
		}
	}
	res2 := make([]string, c)
	r := 0
	for _, z := range res {
		if z != "" {
			res2[r] = z
			r++
		}
	}
	return res2
}

/*func main() {
	str := "Hello how are you?"
	fmt.Println(SplitWhiteSpaces(str))
}*/
