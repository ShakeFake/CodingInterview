package utils

import "fmt"

func SwapArrayWithTwoIndex(list []int, index1 int, index2 int) {
	swap := list[index1]
	list[index1] = list[index2]
	list[index2] = swap
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func LinkPrint(data *ListNode) {
	for {
		if data == nil {
			break
		}

		fmt.Println(data.Val)
		if data.Next != nil {
			data = data.Next
		} else {
			break
		}
	}
}
