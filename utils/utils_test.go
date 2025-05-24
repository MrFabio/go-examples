package utils

import (
	"testing"
)

func TestIntArrToStrArr(t *testing.T) {
	arr := []int{1, 2, 3}
	strArr := IntArrToStrArr(arr)
	if len(strArr) != len(arr) {
		t.Errorf("Expected %d, got %d", len(arr), len(strArr))
	}
}
