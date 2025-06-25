package leetcode

import (
	"fmt"
	"testing"
)

/*
  问题：给定一个未经排序的不重复的分数数组，和一个可能的最大值。返回一个排过序的和符合该限制的数组。
  例子:    输入: []int{11, 12, 13, 14, 15}, Max: 15
           输出: Output: []int{15, 14, 13, 12, 11}
           输入: []int{11, 12, 98, 14, 15}, Max: 100
           输出: Output: []int{98, 15, 14, 12, 11}
  方法：使用最大值构造数组，然后将分数依次填入，最后遍历该数组进行返回。
		应该有更多的限制条件。不然，直接对该数组排序即可，添加一个排序规则。
*/

func TestMaxScopeSort(t *testing.T) {
	maxScopeSorts := []struct {
		Input  []int
		Max    int
		Output []int
	}{
		{Input: []int{11, 12, 13, 14, 15}, Max: 15, Output: []int{15, 14, 13, 12, 11}},
		{Input: []int{11, 12, 98, 14, 15}, Max: 100, Output: []int{98, 15, 14, 12, 11}},
	}

	for _, maxScopeSort := range maxScopeSorts {
		result := MaxScopeSort(maxScopeSort.Input, maxScopeSort.Max)
		fmt.Println(result)
	}
}

func MaxScopeSort(dates []int, maxScope int) []int {
	// 注意，此处不能使用map，map的len是指key的数目。
	temp := make([]int, maxScope+1)
	out := []int{}

	for _, v := range dates {
		temp[v]++
	}

	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] >= 1 {
			out = append(out, i)
		}
	}
	return out
}
