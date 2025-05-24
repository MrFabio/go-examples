package functions

import (
	"math/rand"
)

// RandExamples demonstrates random number generation in Go
// including seeding, basic random numbers, and random ranges
func RandExamples() {
	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	pl("Random int:", rand.Intn(10))
	pl("Random float:", rand.Float64())
	pl("Random float between 0 and 10:", rand.Float64()*10)
	// shuffle
	arr := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	pl("Shuffled array:", arr)
}
