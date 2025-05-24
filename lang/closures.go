package lang

func useFunction(f func(int, int) int, x, y int) {
	pl("Result: ", f(x, y))
}

func ClosuresExamples() {
	intSum := func(x, y int) int {
		return x + y
	}

	pl(intSum(1, 2))

	samp1 := 1
	changeVar := func() {
		samp1 += 1
	}

	changeVar()
	pl("samp1: ", samp1)

	useFunction(intSum, 5, 2)
}
