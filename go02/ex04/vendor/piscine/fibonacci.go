package piscine

func fib(n int, a int, b int) int {
	if n == 0 {
		return b
	}
	return fib(n-1, b, a+b)
}

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	return fib(index, 1, 0)
}
