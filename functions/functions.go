package functions

import (
	utils "go-examples/utils"
)

// FunctionsExamples demonstrates various function features in Go
// including variadic parameters, multiple return values, and closures
func FunctionsExamples() {
	// Variadic function example
	utils.Pl(sumNums(1, 2, 3, 4, 5))

	// Multiple return values
	quotient, remainder := divide(10, 3)
	utils.Pl(quotient, remainder)

	// Named return values
	utils.Pl(divideNamed(10, 3))

	// Closure example
	counter := createCounter()
	utils.Pl(counter()) // 1
	utils.Pl(counter()) // 2
	utils.Pl(counter()) // 3
}

// sum calculates the sum of any number of integers
func sumNums(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// divide performs integer division and returns quotient and remainder
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// divideNamed demonstrates named return values
func divideNamed(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// createCounter returns a closure that maintains its own counter
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
