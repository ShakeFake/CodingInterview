package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/*
	问题：合并两个已经排过序的数组
	例子：
		输入: {1, 2, 4, 6} {3, 4, 5, 7}
		输出: {1, 2, 3, 4, 5, 6, 7}
		输入: {1, 2, 3, 4} {2, 3, 4, 5}
		输出: {1, 2, 3, 4, 5}
	方法：对两个数组进行遍历，依次拼接，最后将不足的数组拼接至结果后面。
*/

func TestMergeSortedArrays(t *testing.T) {
	testMergeSortedArrays := []struct {
		Input  [][]int
		OutPut []int
	}{
		{Input: [][]int{{1, 2, 4, 6}, {3, 4, 5, 7}}, OutPut: []int{1, 2, 3, 4, 5, 6, 7}},
		{Input: [][]int{{1, 2, 3, 4}, {2, 3, 4, 5}}, OutPut: []int{1, 2, 3, 4, 5}},
	}

	for _, testMergeSortedArray := range testMergeSortedArrays {
		mergeResult := MergeSortedArray(testMergeSortedArray.Input)
		fmt.Println(mergeResult)
	}
}

// 这样会有重复的，相同的值不应该塞进数组当中。
func MergeSortedArraySimple(dates [][]int) (result []int) {
	result = []int{}
	for i := 0; i < len(dates); i++ {
		for j := 0; j < len(dates[i]); j++ {
			result = append(result, dates[i][j])
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return
}

func MergeSortedArray(dates [][]int) []int {

	first := dates[0]
	second := dates[1]

	result := []int{}
	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			result = append(result, first[i])
			i++
			continue
		}

		if first[i] > second[j] {
			result = append(result, second[j])
			j++
			continue
		}

		result = append(result, second[j])
		i++
		j++
	}

	// 注意，i，j从0开始，每次计数加1，所以此处比较时，不需要再减一。
	if i == len(first) {
		result = append(result, second[j:]...)
	}

	if j == len(second) {
		result = append(result, first[i:]...)
	}
	return result
}
