package datatypes

import (
	"fmt"
	"unicode/utf8"
)

// RuneExamples demonstrates various operations with Go runes
// including basic declarations, string operations, and Unicode handling
func RuneExamples() {
	// Basic rune declaration and properties
	var r rune = 'A'
	fmt.Printf("Rune: %c\n", r)
	fmt.Printf("Unicode value: %d\n", r)
	fmt.Printf("Type: %T\n", r)

	// String and runes with Unicode characters
	str := "Hello, 世界"
	fmt.Printf("\nString: %s\n", str)
	fmt.Printf("Length in bytes: %d\n", len(str))
	fmt.Printf("Length in runes: %d\n", utf8.RuneCountInString(str))

	// Iterating over runes in a string
	fmt.Println("\nIterating over runes:")
	for i, r := range str {
		fmt.Printf("Position %d: %c (Unicode: %d)\n", i, r, r)
	}

	// Converting between strings and runes
	runes := []rune(str)
	fmt.Printf("\nRune slice length: %d\n", len(runes))

	// Special Unicode characters
	specialChars := "☺★♠♣"
	fmt.Println("\nSpecial characters:")
	for _, r := range specialChars {
		fmt.Printf("Character: %c, Unicode: %U\n", r, r)
	}

	// Rune literals and escape sequences
	newline := '\n'
	tab := '\t'
	fmt.Printf("\nEscape sequences:\nNewline:%c(Unicode: %d)\nTab:%c(Unicode: %d)\n",
		newline, newline, tab, tab)
}
