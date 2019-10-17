package piscine

import "fmt" 

func PointOne(n *int) { // в функции PointOne n - указывает на int
	*n = 1

}
//Используя указатель (*int) в функции PointOne, мы изменяем значение оригинальной переменной.