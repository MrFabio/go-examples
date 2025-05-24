package datatypes

import "fmt"

// ArraysExamples demonstrates various array operations in Go
// including declaration, initialization, and manipulation
func ArraysExamples() {
	// Array declaration and initialization
	var arr1 [5]int
	arr1[0] = 1
	pl(arr1)

	// Array literal initialization
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl(arr2[0])
	pl(len(arr2))

	// Array iteration using for loop
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	// Array iteration using range
	for i, v := range arr1 {
		pl(i, v)
	}

	// Multi-dimensional array
	arr3 := [2][2]int{{1, 2}, {3, 4}}
	pl(arr3[0][0])
	pl(len(arr3))

	// Multi-dimensional array iteration
	for i := 0; i < len(arr3); i++ {
		pl(arr3[i])
	}

	// String to rune array conversion
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		pl(v)
	}

	// Array operations
	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)
	fmt.Println("Array 3:", arr3)
	fmt.Println("Length of arr2:", len(arr2))
	fmt.Println("First element of arr2:", arr2[0])
	fmt.Println("Last element of arr2:", arr2[len(arr2)-1])

	// Array iteration
	fmt.Println("\nIterating over arr2:")
	for i, v := range arr2 {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	// Multi-dimensional array
	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}
	fmt.Println("\nMatrix:", matrix)
}
