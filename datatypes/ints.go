package datatypes

import (
	"fmt"
	"reflect"
	"strconv"
)

// IntsExamples demonstrates various integer operations in Go
// including type conversion, parsing, and formatting
func IntsExamples() {
	// Basic variable declarations
	var name string = "World"
	var v1, v2 int = 1, 2
	var v3 = name
	v4 := 2.4

	// Type conversion and reflection
	v1 = v2
	pl(reflect.TypeOf(v4))
	pl(int(v4))

	// String to int conversion
	v3 = "4"
	pl(strconv.Atoi(v3))
	pl(strconv.Itoa(v1))

	// Float parsing
	if cV8, err := strconv.ParseFloat("3.14", 64); err != nil {
		pl("Error parsing float", err)
	} else {
		cv9 := fmt.Sprintf("%f", cV8)
		pl("Float value is", cV8, "String value is", cv9)
	}
}
