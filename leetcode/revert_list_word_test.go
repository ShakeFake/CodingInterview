package leetcode

import (
	"fmt"
	"testing"
)

/*
	问题：给定一句话，将这句话所有单词翻转。
	例子：
		输入：[]string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"}
		输出：[]string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"}
	方法：第一次翻转翻转全部，然后在将每个单词额外翻转即可。
*/

func TestRevertList(t *testing.T) {
	testRevertDatas := []struct {
		Input  []string
		Output []string
	}{
		// 因为你要区分单词的反转，因此必须要有信号区分单词。
		{[]string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"},
			[]string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"}},
	}

	for _, testRevertData := range testRevertDatas {
		resultRevertListWord := RevertListWord(testRevertData.Input)
		fmt.Println(resultRevertListWord)
	}

}

func RevertListWord(dates []string) []string {
	revertResult := RevertListString(dates)
	start := 0
	for k, v := range revertResult {
		if v == "" {
			RevertListString(revertResult[start:k])
			start = k + 1
		}
	}
	// 此处使用 slice切片 的属性，修改值会对原值造成影响。
	RevertListString(revertResult[start:len(revertResult)])
	return revertResult
}
