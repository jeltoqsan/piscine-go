package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
} 

// i = i + 1 - можно и так записать 	
// i - переменная, ее значение 'a', если 'а' условно меньше или равно 10, то выводится букву 'b' и дойдя до z равно z -цикл останавливается и срабатывает PrintRune; \n - newline - перевод на новую строку
// 2 способ: func main () { //	fmt.Println ("abcdefghijklmnopqrstuvwxyz") } 
