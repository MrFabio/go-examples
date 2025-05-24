package utils

import (
	"fmt"
	"strconv"
)

// Pl is a convenience variable that aliases fmt.Println
var Pl = fmt.Println

// IntArrToStrArr converts a slice of integers to a slice of strings
func IntArrToStrArr(arr []int) []string {
	var strArr = []string{}

	for _, v := range arr {
		strArr = append(strArr, strconv.Itoa(v))
	}

	return strArr
}
