package main

import . "github.com/01-edu/z01"

func Rot14(str string) string {
	nb := rune(14)
	result := ""
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			if arr[i] >= 'm' {
				arr[i] = arr[i] - (nb - 2)
			} else {
				arr[i] = arr[i] + nb
			}
		} else if arr[i] >= 'A' && arr[i] <= 'Z' {
			if arr[i] >= 'M' {
				arr[i] = arr[i] - (nb - 2)
			} else {
				arr[i] = arr[i] + nb
			}
		}
		result += string(arr[i])
	}
	return result
}
func main() {
	result := Rot14("Hello How are You")
	arrayRune := []rune(result)

	for _, s := range arrayRune {
		PrintRune(s)
	}
	PrintRune('\n')
}
