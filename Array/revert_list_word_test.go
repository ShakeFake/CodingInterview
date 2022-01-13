package test

import (
	"fmt"
	"testing"
)

func TestRevertList(t *testing.T) {
	testRevertDatas := []struct {
		Input  []string
		Output []string
	}{
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
	RevertListString(revertResult[start:len(revertResult)])
	return revertResult
}
