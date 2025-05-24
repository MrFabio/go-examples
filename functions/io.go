package functions

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

const fileName = "data.txt"

// IoExamples demonstrates file I/O operations in Go
// including file creation, writing, reading, and appending
func IoExamples() {
	// Create a new file
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Convert integer array to string array
	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}

	// Write strings to file
	for _, s := range sPrimeArr {
		_, err = f.WriteString(s + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Open file for reading
	f, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read and print file contents
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime:", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	// Append data to existing file
	_, err = os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("File does not exist")
	} else {
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}

// hello demonstrates basic console I/O
func hello() {
	pl("Hello, World!")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello, " + name)
	} else {
		log.Fatal("Error reading input")
	}
}
