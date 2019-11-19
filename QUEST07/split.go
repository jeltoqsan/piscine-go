package main

import "fmt"

func Split(str, charset string) []string {
	counter := 0
	for range charset {
		counter++
	}
	runes := []rune(str)
	xCounter := 0
	for range runes {
		xCounter++
	}

	chRunes := []rune(charset)
	char := 0
	checker := make([]rune, counter)
	for i := 0; i < xCounter-counter; i++ {
		for j := 0; j < counter; j++ {
			checker[j] = runes[i+j]
		}
		res := true
		for k := 0; k < counter; k++ {
			if checker[k] != chRunes[k] {
				res = false
			}
		}
		if res {
			char++
		}
	}
	result := make([]string, char+1)
	dot := 0
	x := 0
	for i := 0; i < xCounter-counter; i++ {
		for j := 0; j < counter; j++ {
			checker[j] = runes[i+j]
		}
		res := true
		for k := 0; k < counter; k++ {
			if checker[k] != chRunes[k] {
				res = false
			}
		}
		if res {
			result[x] = string(str[dot:i])
			x++
			dot = i + counter
		}
	}
	result[x] = string(str[dot:])
	return result
}
func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}
