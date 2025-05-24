package lang

import (
	"fmt"
	"log"
)

func AccessorsExamples() {
	date := Date{}
	err := date.SetDay(25)
	if err != nil {
		log.Fatal("Error setting day:", err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal("Error setting month:", err)
	}
	err = date.SetYear(2025)
	if err != nil {
		log.Fatal("Error setting year:", err)
	}
	fmt.Printf("main date: %d/%d/%d\n", date.Day(), date.Month(), date.Year())
}
