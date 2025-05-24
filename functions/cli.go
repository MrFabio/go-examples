package functions

import (
	"fmt"
	"os"
	"strconv"
)

// CliExamples demonstrates command-line interface features in Go
// including argument parsing and processing
func CliExamples() {
	// Print all command-line arguments
	fmt.Println(os.Args)

	// Get arguments excluding program name
	args := os.Args[1:]
	var iArgs = []int{}

	// Convert string arguments to integers
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, i)
	}

	// Find maximum value in arguments
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max:", max)
}
