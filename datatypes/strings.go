package datatypes

import (
	"strings"
)

// StringsExamples demonstrates various string operations in Go
// including concatenation, formatting, and manipulation
func StringsExamples() {
	// Basic string operations
	var text string = "Hello, World!"
	replacer := strings.NewReplacer("Hello", "Goodbye", "World", "Cruel World")
	text = replacer.Replace(text)
	pl(text)

	// String length and searching
	pl("Length of text is", len(text))
	pl("Contains World?", strings.Contains(text, "World"))
	pl("Index of World is", strings.Index(text, "World"))

	// String manipulation
	pl("Replace o with 0:", strings.Replace(text, "o", "0", -1))
	pl("Trim whitespaces:", strings.TrimSpace("  \n\"Hello, World!  "))
	pl("Split by comma:", strings.Split("Hello, World!", ","))
	pl("Repeat 3 times:", strings.Repeat("Hello, World! ", 3))
	pl("Lowercase:", strings.ToLower("Hello, World!"))
	pl("Uppercase:", strings.ToUpper("Hello, World!"))
	pl("Title:", strings.Title("hello, world!"))
	pl("Trim prefix:", strings.TrimPrefix("Hello, World!", "Hello"))
	pl("Trim suffix:", strings.TrimSuffix("Hello, World!", "World!"))
	pl("Trim:", strings.Trim("  \n\"Hello, World!  ", " \n\""))
	pl("Fields:", strings.Fields("Hello, World!"))
}
