package Hash

import (
	"fmt"
	"testing"
)

/*
	given a string, judge if it is a permutation palindrome
	example:
		Input: "ivciv"
		Output: true

		Input: "abcba"
		Output: true

		Input: "abcde"
		Output: true
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
