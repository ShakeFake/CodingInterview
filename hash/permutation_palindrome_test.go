package hash

import (
	"fmt"
	"testing"
)

/*
  问题：给定一个字符串，判断该字符串的排列是否为回文。
  例子:	输入: "ivciv"
		输出: true
		输入: "abcba"
		输出: true
  方法：1：如果有亦或运算符会好一些。
		2：使用一个map存储每次的字母，相同删除，不同去除，最后判断map长度即可。
*/

func TestPermutationPalindrome(t *testing.T) {

	testPermutationPalindromes := []struct {
		Input  string
		Output bool
	}{
		{Input: "ivciv", Output: true},
		{Input: "abcba", Output: true},
		{Input: "abcde", Output: false},
	}
	for _, testPermutationPalindrome := range testPermutationPalindromes {
		result := PermutationPalindrome(testPermutationPalindrome.Input)
		fmt.Println(result)
	}
}

func PermutationPalindrome(date string) bool {
	index := map[int32]int{}

	for _, item := range date {
		if index[item] > 0 {
			delete(index, item)
		} else {
			index[item] = 1
		}
	}
	// 不敢相信，下述两行代码逻辑相同。
	return len(index) <= 1
	//if len(index) > 1 {
	//	return false
	//} else {
	//	return true
	//}
}
