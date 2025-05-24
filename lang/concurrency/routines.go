package concurrency

import "time"

func printTo15() {
	for i := 0; i < 15; i++ {
		pl("F1: ", i)
	}
}

func printTo10() {
	for i := 0; i < 10; i++ {
		pl("F2: ", i)
	}
}

func RoutinesExamples() {
	go printTo15()
	go printTo10()
	// wait for the routines to finish
	// otherwise the main function will exit
	// and the routines will not be executed
	time.Sleep(1 * time.Second)
}
