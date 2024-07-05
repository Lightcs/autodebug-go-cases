package calculate

func Fibonacci(n int) int {
	return Fibonacci(n-1) + Fibonacci(n-2)
}
