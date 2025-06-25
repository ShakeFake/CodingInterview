package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	result := LongestPalindrome("babad")
	fmt.Println(result)

}

func LongestPalindrome(s string) string {
	// 使用动态规划的方法去解题

	if len(s) < 2 {
		return s
	}

	length := len(s)
	begin := 0
	maxlength := 1

	store := make([][]bool, length)
	for i := 0; i < length; i++ {
		store[i] = make([]bool, length)
		store[i][i] = true
	}

	for L := 2; L <= length; L++ {
		for i := 0; i < length; i++ {
			j := L + i - 1

			if j >= length {
				break
			}

			if s[i] == s[j] {
				if j-i < 3 {
					store[i][j] = true
				} else {
					store[i][j] = store[i+1][j-1]
				}
			} else {
				store[i][j] = false
			}

			if store[i][j] && j-i+1 > maxlength {
				maxlength = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxlength]
}
