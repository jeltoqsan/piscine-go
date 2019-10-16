package main 

import "github.com/01-edu/z01"

func main() {
	for latin := 'a'; latin<='z'; latin = latin + 1 { 
		z01.PrintRune(latin) 
	}
	z01.PrintRune('\n') 
} 
//latin++ - можно и так записать 	
// latin - переменная, ее значение 'a', если 'а' условно меньше или равно 10, то выводится букву 'b' и дойдя до z равно z -цикл останавливается и срабатывает PrintRune; \n - newline - перевод на новую строку
//  2 способ: func main () { 
//	fmt.Println ("abcdefghijklmnopqrstuvwxyz") } 