package functions

import (
	"fmt"
	utils "go-examples/utils"
)

var pl = utils.Pl

// CallsExamples demonstrates various function features in Go
// including variadic parameters, multiple return values, and closures
func CallsExamples() {
	// Basic function call with single return value
	pl(sum(1, 2))

	// Function with multiple return values
	r1, r2 := sum2(1, 2)
	pl(r1, r2)

	// Variadic function example
	pl(getSums(1, 2, 3, 4, 5))

	// Function with slice parameter
	pl(getArraySum([]int{1, 2, 3, 4, 5}))

	// Value vs pointer parameter demonstration
	f3 := 3
	pl("f3 before", f3)
	changeValue(f3)
	pl("f3 after ", f3)
	changeValue2(&f3)
	pl("f3 after2", f3)

	// Array pointer parameter
	plArray := [4]int{1, 2, 3, 4}
	dblArray(&plArray)
	pl(plArray)

	// Variadic function with float64
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))
}

// getAverage calculates the average of a variable number of float64 values
func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, num := range nums {
		sum += num
	}
	return sum / numSize
}

// dblArray doubles each element in the given array using pointer
func dblArray(arr *[4]int) {
	for i := range arr {
		arr[i] *= 2
	}
}

// sum returns the sum of two integers
func sum(a int, b int) int {
	return a + b
}

// sum2 returns both the sum and difference of two integers
func sum2(a int, b int) (int, int) {
	return a + b, a - b
}

// getQuotient performs division and handles division by zero error
func getQuotient(a float64, b float64) (ans float64, err error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// getSums calculates the sum of a variable number of integers
func getSums(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	// can omit return var
	return
}

// getArraySum calculates the sum of integers in a slice
func getArraySum(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}
	return
}

// changeValue demonstrates pass by value
func changeValue(f3 int) int {
	f3 += 1
	return f3
}

// changeValue2 demonstrates pass by reference using pointer
func changeValue2(f3 *int) {
	*f3 += 1
}
