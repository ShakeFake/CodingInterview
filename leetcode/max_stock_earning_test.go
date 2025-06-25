package leetcode

import (
	"fmt"
	"testing"
)

/*
  问题：给定一个股票的价格，整数，返回该股票的最大收益值。
  例子:
		输入: []int{6, 10, 14, 20, 5}
		输出: 15
  方法：找到最大值和最小值，然后比较输出即可。或者记录最大值，最小值的顺序
*/

func TestMaxStockEarning(t *testing.T) {
	maxStockEarnings := []struct {
		INPUT  []int
		OUTPUT int
	}{
		{INPUT: []int{6, 10, 14, 20, 5}, OUTPUT: 15},
		{INPUT: []int{50, 10, 25, 20, 27}, OUTPUT: 40},
	}
	for _, maxStockEarning := range maxStockEarnings {
		result := MaxStockEarning(maxStockEarning.INPUT)
		fmt.Println(result)
	}
}

func MaxStockEarning(dates []int) int {
	minPrice := 0
	maxPrice := 0
	for _, v := range dates {
		if minPrice == maxPrice && maxPrice == 0 {
			minPrice = v
			maxPrice = v
			continue
		}
		if v < minPrice {
			minPrice = v
		}
		if v > maxPrice {
			maxPrice = v
		}
	}
	return maxPrice - minPrice
}
