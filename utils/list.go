package utils

func SwapArrayWithTwoIndex(list []int, index1 int, index2 int) {
	swap := list[index1]
	list[index1] = list[index2]
	list[index2] = swap
}
