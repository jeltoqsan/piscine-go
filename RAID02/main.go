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

func solveSudoku(table *[9][9]int) bool {
	if cnt > 1 && !ok {
		return false
	}
	if !hasEmptyCell(table) {
		cnt++

		if ok {
			return !hasEmptyCell(table)
		}
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				ok2 := false
				for candidate := 9; candidate >= 1; candidate-- {
					table[i][j] = candidate

					if isTableValid(table) {
						if solveSudoku(table) {
							if !ok {
								ok2 = true
							} else {
								return true
							}
						}
						table[i][j] = 0
					} else {
						table[i][j] = 0
					}
				}
				if ok2 {
					return true
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(table *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isTableValid(table *[9][9]int) bool {

	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[table[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[table[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

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

func hasDuplicates(counter [10]int) bool {
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

func printTable(table [9][9]int) {
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

func parseInput(input string) [9][9]int {
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

func atoi(s string) int {
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
