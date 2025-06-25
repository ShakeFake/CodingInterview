package leetcode

import (
	"codingInterview/utils"
	"testing"
)

/*
  问题：计算两个链表的数字相加
  例子:	输入: l1 = [2, 4, 3]  l2=[5, 6, 4]
		输出: [7, 0, 8]
  方法：指定分隔符，分割统计即可。
*/

func TestAddTwoNumbers(t *testing.T) {
	testAddTwoNumbers := []utils.ListNode{
		{
			Val: 2,
			Next: &utils.ListNode{
				Val: 4,
				Next: &utils.ListNode{
					Val: 3,
				},
			},
		},
		{
			Val: 5,
			Next: &utils.ListNode{
				Val: 6,
				Next: &utils.ListNode{
					Val: 4,
				},
			},
		},
	}

	result := addTwoNumbers(&testAddTwoNumbers[0], &testAddTwoNumbers[1])

	utils.LinkPrint(result)
}

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	plus := 0
	listArray := []utils.ListNode{}
	listReturn := utils.ListNode{}

	for l1 != nil && l2 != nil {
		addNumber := l1.Val + l2.Val + plus
		plus = 0
		l1 = l1.Next
		l2 = l2.Next
		if addNumber >= 10 {
			plus = 1
			addNumber = addNumber % 10
		}
		listArray = append(listArray, utils.ListNode{addNumber, nil})
	}

	for l1 != nil {
		addNumber := l1.Val + plus
		plus = 0
		l1 = l1.Next
		if addNumber >= 10 {
			plus = 1
			addNumber = addNumber % 10
		}
		listArray = append(listArray, utils.ListNode{addNumber, nil})
	}

	for l2 != nil {
		addNumber := l2.Val + plus
		plus = 0
		l2 = l2.Next
		if addNumber >= 10 {
			plus = 1
			addNumber = addNumber % 10
		}
		listArray = append(listArray, utils.ListNode{addNumber, nil})
	}

	if plus > 0 {
		listArray = append(listArray, utils.ListNode{plus, nil})
	}

	for i := len(listArray) - 1; i >= 0; i-- {
		listReturn.Val = listArray[i].Val
		middle := listReturn
		listReturn.Next = &middle
	}

	return listReturn.Next
}
