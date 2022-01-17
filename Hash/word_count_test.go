package Hash

import (
	"fmt"
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	testWordCounts := []struct {
		Input  string
		Output map[string]int
	}{
		{"i am a boy", map[string]int{"i": 1, "am": 1, "a": 1, "boy": 1}},
		{"i am am a boy boy", map[string]int{"i": 1, "am": 2, "a": 1, "boy": 2}},
	}

	for _, testWordCount := range testWordCounts {
		result := WordCount(testWordCount.Input)
		fmt.Println(result)
	}

}

func WordCount(date string) (result map[string]int) {
	result = map[string]int{}
	for _, item := range strings.Split(date, " ") {
		result[item]++
	}
	return
}
