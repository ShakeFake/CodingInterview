package Hash

import (
	"fmt"
	"testing"
)

func TestDoubleNumAdd(t *testing.T) {
	testDoubleNumAddDates := []struct {
		Input  []int
		Output int
	}{
		{Input: []int{1, 2, 3, 4, 5}, Output: 8},
		{Input: []int{3, 5, 7, 11, 13}, Output: 16},
		{Input: []int{3, 5, 7, 11, 13}, Output: 25},
	}

	for _, testDoubleNumAddDate := range testDoubleNumAddDates {
		result := DoubleNumAdd(testDoubleNumAddDate.Input, testDoubleNumAddDate.Output)
		fmt.Println(result)
	}
}

func DoubleNumAdd(dates []int, expect int) bool {
	index := map[int]bool{}

	for _, v := range dates {
		askFor := expect - v
		if index[askFor] {
			return true
		} else {
			// 此处存储的是上一个值是否存在。不同于找到所有可行性。
			index[v] = true
		}
	}
	return false
}
