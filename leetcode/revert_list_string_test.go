package leetcode

import (
	"fmt"
	"testing"
)

/*
	问题：给定一个字符串，翻转它。
	例子：
		输入：[]string{"a", "b", "c", "d"}
		输出：[]string{"d", "c", "b", "a"}
	方法：翻转数组即可
*/

func TestRevertListString(t *testing.T) {
	testRevertDatas := []struct {
		Input  []string
		OutPut []string
	}{
		{Input: []string{"a"}, OutPut: []string{"a"}},
		{Input: []string{"a", "b"}, OutPut: []string{"b", "a"}},
		{Input: []string{"a", "b", "c"}, OutPut: []string{"c", "b", "a"}},
		{Input: []string{"a", "b", "c", "d"}, OutPut: []string{"d", "c", "b", "a"}},
		{Input: []string{"a", "b", "c", "d", "e"}, OutPut: []string{"e", "d", "c", "b", "a"}},
	}

	for _, v := range testRevertDatas {
		revertResult := RevertListString(v.Input)
		fmt.Println(revertResult)
	}

}

func RevertListString(dates []string) []string {

	length := len(dates) - 1

	// 此处可以封装公共模块替换。或者直接减半处理。
	for i := 0; i < length/2+1; i++ {
		temp := dates[i]
		dates[i] = dates[length-i]
		dates[length-i] = temp
	}

	return dates
}
