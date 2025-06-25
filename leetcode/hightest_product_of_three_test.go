package leetcode

import (
	. "codingInterview/utils"
	"fmt"
	"sort"
	"testing"
)

/*
  问题：给定一个整数数组，计算该整数数组当中乘积最大的三个值的结果
  例子:
		输入: []int{-10, -10, 3, 2, 1}
		输出: 300
  方法：我们将数组进行排序，从小到大，然后计算 -1 -2 -3 和 -1 0 1 之间的最大值即可。
		详细论证算法：https://leetcode-cn.com/problems/maximum-product-of-three-numbers/solution/tan-xin-zheng-ming-onshi-jian-by-boille-9kxv/
*/

func TestHighestProduceThree(t *testing.T) {
	highestProduceThrees := []struct {
		INPUT  []int
		OUTPUT int
	}{
		{INPUT: []int{-10, -10, 3, 2, 1}, OUTPUT: 300},
		{INPUT: []int{-50, 6, 3, 2, 1}, OUTPUT: 36},
	}
	for _, highestProduceThree := range highestProduceThrees {
		result := HighestProduceThree(highestProduceThree.INPUT)
		fmt.Println(result)
	}
}

func HighestProduceThree(dates []int) int {
	sort.Slice(dates, func(i, j int) bool {
		// 从小到大
		return dates[i] <= dates[j]
	})
	length := len(dates)
	// 此处采用了一个算法推论。
	return MaxOfTwoNumber(dates[length-1]*dates[length-2]*dates[length-3], dates[0]*dates[1]*dates[length-1])
}
