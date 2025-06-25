package leetcode

import (
	"fmt"
	"testing"
)

/*
  问题：无重复字符的最长子串
  例子:	输入: "abcabcbb"
		输出: 3
  方法：使用hash
*/

func TestLengthOfLongestSubstring(t *testing.T) {
	testDatas := []string{
		"pwwkew",
	}

	for _, testData := range testDatas {
		result := LengthOfLongestSubstring(testData)
		fmt.Println(result)
	}

}

func LengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	// 记录统计过多少值
	first := 0

	// 记录最大长度
	maxLength := 0

	maxString := make(map[string]bool, 0)
	for _, v := range s {
		v := string(v)
		if maxString[v] {
			for {
				delete(maxString, s[first:first+1])
				// 从后往前删除
				first++
				if !maxString[v] {
					break
				}
			}

		}
		maxString[v] = true
		if len(maxString) > maxLength {
			maxLength = len(maxString)
		}
	}
	return maxLength
}
