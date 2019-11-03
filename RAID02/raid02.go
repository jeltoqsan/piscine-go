package main

import (
	"fmt"
	"os"
)

var cnt int
var ok bool

func main() {
	if len(os.Args[1:]) != 9 {
		fmt.Println("Error")
		return
	}

	for _, arg := range os.Args[1:] {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}

		for _, ch := range arg {
			if ch < '0' || ch > '9' {
				if ch != '.' {
					fmt.Println("Error")
					return
				}
			}
		}
	}

	arr := os.Args[1:]
	StrFromArr := ""
	for _, valueArr := range arr {
		for i := 0; i <= 8; i++ {
			if valueArr[i] == '.' {
				StrFromArr = StrFromArr + "0"
			} else {
				StrFromArr = StrFromArr + string(valueArr[i])
			}

		}
	}

	table := parseInput(StrFromArr)

	if solveSudoku(&table) {
		if cnt == 1 {
			ok = true
			solveSudoku(&table)
			printTable(table)
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}

}

func solveSudoku(table *[9][9]int) bool { // запускаем наше решение , в виде рекурсии,
	if cnt > 1 && !ok {
		return false
	}
	if !hasEmptyCell(table) { // проверяет если у нас, свободная ячейка в нашем судоку,
		cnt++
		//printTable(table)

		if ok {
			return !hasEmptyCell(table) // если у нас все заполнено , то мы возрващаем тру
		}
		return true
	}
	for i := 0; i < 9; i++ { // перебираем каждую ячейку нашего судоку, при нахождении незаполннной ячейки, пытаемся подставить определенного кандидата на эту ячейку, путем проверки через функцию из тэйбл валлет, при ее успешном проходжении запускаем, рекурскию, с обновленным значением тэйюл в определнной ячейке
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				ok2 := false
				for candidate := 9; candidate >= 1; candidate-- {
					table[i][j] = candidate

					if isTableValid(table) {
						if solveSudoku(table) {
							if !ok {
								ok2 = true // при подставлении данного кандидата и при дальнейшем его прохождении, то есть захождения в рекурскию солв судоку, мы получили положит ответ, то есть возможность построения судоку с данным кандидатом
							} else {
								return true
							}
						}
						table[i][j] = 0 // из за того что у нас перебор, мы долджны освободить эту ячекйку для дашльшенго использования
					} else {
						table[i][j] = 0 // путем подставления данного кандидата в эту ячейку наше изначальое судоку, стало невалидный, то есть нерешаемым, мы отсекаем этот вариант
					}
				}
				if ok2 {
					return true
				}
				return false // из за при подставлении кандидата и не смогли решить судоку, при переборе всех кандидатов, то есть несуществут такой цифры которой, что при подставлении наше судоку становится решаемымы изщ за это возращаем фолс
			}
		}
	}
	return false
}

func hasEmptyCell(table *[9][9]int) bool { // проверка на пустые строки
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isTableValid(table *[9][9]int) bool { // являются ли судоку решаемым.только тогда когда все цифры в строка столцах и квадратитак уникальные, то есть различаются друг от друга

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[table[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[table[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[table[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool { // проверка на уникальность, цифр на неповторяемость
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func printTable(table [9][9]int) { // функция отвеяает за вывод нашей таблицы
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if col != 8 {
				fmt.Printf("%d ", table[row][col])
			} else {
				fmt.Printf("%d", table[row][col])
			}
		}
		fmt.Println()
	}
}

func parseInput(input string) [9][9]int { // это функция, которая преедлывает нашу изначальную строку в судоку со значения цифр где 0 отсутствие элемента
	table := [9][9]int{}
	i := 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			i1 := atoi(string(input[i]))
			i++
			table[row][col] = i1
		}
	}
	return table
}

func atoi(s string) int { // переводит стринг в числа чтобы использдовать его в нашем двухмерном массиве тэйбл
	if len(s) == 0 {
		return 0
	}

	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if len(s) < 1 {
			return 0
		}
	}

	nm := 0

	for _, ch := range s {
		if !containsIn0to9(ch) {
			return 0
		}
		nm = nm*10 + charToInt(ch)
	}

	if s0[0] == '-' {
		nm *= -1
	}
	return nm
}

func containsIn0to9(ch rune) bool {
	for i := '0'; i <= '9'; i++ {
		if ch == i {
			return true
		}
	}

	return false
}

func charToInt(ch rune) int {
	count := 0
	if ch < 48 && ch > 58 {
		return 0
	}

	for i := '1'; i <= ch; i++ {
		count++
	}

	return count
}
