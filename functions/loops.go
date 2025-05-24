package functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// LoopsExamples demonstrates various loop constructs in Go
// including for loops, while loops, and range loops
func LoopsExamples() {
	// for loop
	for i := 0; i < 10; i++ {
		pl(i)
	}

	// for range
	for i := range 10 {
		pl(i)
	}

	// while loop
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	fX = 10
	// random integer
	randInt := rand.Intn(10)
	fmt.Println("Guess the number")
	for {
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			pl(err)
		}
		guess = strings.TrimSpace(guess)
		guessInt, err := strconv.Atoi(guess)
		if err != nil {
			pl(err)
		}
		if guessInt < randInt {
			pl("Guess Higher")
		} else if guessInt > randInt {
			pl("Guess Lower")
		} else {
			pl("You guessed the number!")
			break
		}
	}

	// arrays
	arr := [5]int{1, 2, 3, 4, 5}
	pl(arr)

	for _, value := range arr {
		pl(value)
	}
}
