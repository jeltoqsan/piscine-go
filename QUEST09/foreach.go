//package piscine
package main

func ForEach(f func(int), arr []int) {
	for _, r := range arr {
		f(r)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, arr)
}
func PrintNbr(int) {

}
