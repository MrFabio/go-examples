package datatypes

// SlicesExamples demonstrates various slice operations in Go
// including creation, modification, and manipulation
func SlicesExamples() {
	// Create and initialize a string slice
	sl1 := make([]string, 5)
	sl1[0] = "Sporting"
	sl1[1] = "Clube"
	sl1[2] = "de"
	sl1[3] = "Braga"
	sl1[4] = "SAD"
	pl(sl1)

	// Create a slice with length 3 and capacity 5
	names := make([]int, 3, 5)
	pl(names)

	// Create array and slice from it
	oriArray := [5]int{1, 2, 3, 4, 5}
	sliceArray := oriArray[1:3]
	pl(sliceArray)

	// Modify slice and see changes in original array
	sliceArray[0] = 100
	pl(sliceArray)
	pl(oriArray)
	oriArray[2] = 200
	pl(sliceArray)
	pl(oriArray)

	// Append elements to slice
	sliceArray = append(sliceArray, 6, 7, 8, 9, 10)
	pl(sliceArray)

	// Append array elements to slice
	sliceArray = append(sliceArray, oriArray[:]...)
	pl(sliceArray)

	// Create new slice from existing one
	sliceArray = sliceArray[2:]
	pl(sliceArray)
}
