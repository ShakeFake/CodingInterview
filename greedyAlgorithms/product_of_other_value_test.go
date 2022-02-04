package greedyAlgorithms

import (
	"fmt"
	"testing"
)

/*
  问题：给定一个数组，计算该数组除当前位外，其余位的乘积。
  例子:
		输入: []int{2, 3, 4, 5}
		输出: []int{60, 40, 30, 24}
  方法：第一次遍历，计算顺序前置乘积。
		第二次遍历，计算倒叙前置乘积与第一次结果当前值。
*/

func TestProductOtherValue(t *testing.T) {
	productOtherValues := []struct {
		INPUT  []int
		OUTPUT []int
	}{
		{INPUT: []int{2, 3, 4, 5}, OUTPUT: []int{60, 40, 30, 24}},
		{INPUT: []int{1, 7, 3, 4}, OUTPUT: []int{84, 12, 28, 21}},
	}
	for _, productOtherValue := range productOtherValues {
		result := ProductOtherValue(productOtherValue.INPUT)
		fmt.Println(result)
	}
}

func ProductOtherValue(dates []int) []int {
	if len(dates) <= 2 {
		return dates
	}
	out := make([]int, len(dates))
	start1 := 1
	for k, v := range dates {
		if k == 0 {
			out[k] = start1
			start1 = start1 * v
			continue
		}
		out[k] = start1
		start1 = start1 * v
	}

	start2 := 1
	for i := len(dates) - 2; i >= 0; i-- {
		start2 = start2 * dates[i+1]
		out[i] = out[i] * start2
	}
	return out
}
