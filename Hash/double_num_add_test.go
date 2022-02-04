package Hash

import (
	"fmt"
	"testing"
)

/*
  问题：给定一个数组和一个值，判断数组当中是否存在两个数字相加等于该值的。
  例子:    输入: []int{1, 2, 3, 4, 5}, Output: 8
           输出: true
           输入: []int{3, 5, 7, 11, 13}, Output: 16
           输出: true
  方法：使用桶的算法进行处理，遍历数组，查询每一个值对应是否存在
		不存在则由该值减去遍历数字，存储到数组当中。
*/

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
