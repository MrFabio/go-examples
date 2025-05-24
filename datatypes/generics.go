package datatypes

// MyConstraint defines a type constraint for generics
// It allows only int and float64 types
type MyConstraint interface {
	int | float64
}

// getSumGen is a generic function that works with any type
// that satisfies MyConstraint
func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

// GenericsExamples demonstrates the use of generic functions
func GenericsExamples() {
	pl("Generics Examples")

	pl("5 + 4 =", getSumGen(5, 4))
	pl("5.1 + 4.2 =", getSumGen(5.1, 4.2))
}
