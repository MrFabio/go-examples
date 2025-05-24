package lang

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func RecursionExamples() {
	pl("Factorial of 5: ", factorial(5))
}
