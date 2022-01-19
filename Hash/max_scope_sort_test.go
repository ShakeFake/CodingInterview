package Hash

import (
	"fmt"
	"testing"
)

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
