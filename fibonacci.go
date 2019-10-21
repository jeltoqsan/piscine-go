package piscine

func Fibonacci(index int) int {
	if index <= 1 {
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
