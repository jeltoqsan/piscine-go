package piscine

func PointOne(n *int) { // в функции PointOne n - указывает на int
	*n = *n + 1
}

//Используя указатель (*int) в функции PointOne, мы изменяем значение оригинальной переменной.
