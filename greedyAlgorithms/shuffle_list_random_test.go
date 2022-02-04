package greedyAlgorithms

import (
	. "codingInterview/utils"
	"fmt"
	"math/rand"
	"testing"
)

/*
  问题：给定一个数组，以当前位置为基准，进行洗牌。
  例子:
		输入: []int{1,2,3,4,5,6}
		输出: 随机
  方法： 此题意是否不清楚，需要额外思考一下。
*/

func TestShuffleListRandom(t *testing.T) {
	shuffleListRandoms := []struct {
		INPUT []int
	}{
		{INPUT: []int{6, 10, 14, 20, 5}},
	}
	for _, shuffleListRandom := range shuffleListRandoms {
		result := ShuffleListRandom(shuffleListRandom.INPUT)
		fmt.Println(result)
	}
}

func ShuffleListRandom(dates []int) []int {
	if len(dates) <= 1 {
		return dates
	}

	lastIndex := len(dates) - 1
	for k, _ := range dates {
		randomIndex := rand.Intn((lastIndex - k) + k)
		SwapArrayWithTwoIndex(dates[0:5], k, randomIndex)
	}

	return dates
}
