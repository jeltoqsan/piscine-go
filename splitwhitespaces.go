package piscine

/*package main

import "fmt"*/

func SplitWhiteSpaces(str string) []string {
	counter := 0
	for _, r := range str {
		if r == ' ' || r == '\n' || r == '\t' {
			counter++
		}
	}
	res := make([]string, counter+1)
	j := 0
	last := 0
	for i, r := range str {
		if r == ' ' || r == '\n' || r == '\t' {
			if str[last:i] != "" {
				res[j] = str[last:i]
				j++
			}
			last = i + 1
		}
	}
	res[j] = str[last:]
	c := 0
	for _, r := range res {
		if r != "" {
			c++
		}
	}
	res2 := make([]string, c)
	r := 0
	for _, r := range res {
		if r != "" {
			res2[r] = r
			r++
		}
	}
	return res2
}

/*func main() {
	str := "Hello how are you?"
	fmt.Println(SplitWhiteSpaces(str))
}*/
