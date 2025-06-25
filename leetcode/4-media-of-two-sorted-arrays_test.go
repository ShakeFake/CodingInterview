package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {

	result := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println(result)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	allLength := len(nums1) + len(nums2)

	times := allLength / 2
	mulity := allLength % 2

	numberCurrentPre := math.MinInt32
	numberCurrent := math.MinInt32

	// 分别遍历两个数组，有序。
	beforeFirst := 0
	lastFirst := len(nums1)

	beforeSecond := 0
	lastSecond := len(nums2)

	for i := 0; i <= times; i++ {
		numberCurrentPre = numberCurrent
		if beforeFirst < lastFirst && beforeSecond < lastSecond {
			if nums1[beforeFirst] <= nums2[beforeSecond] {
				numberCurrent = nums1[beforeFirst]
				beforeFirst++
			} else {
				numberCurrent = nums2[beforeSecond]
				beforeSecond++
			}
		} else if beforeFirst > lastFirst-1 {
			if numberCurrent <= nums2[beforeSecond] {
				numberCurrent = nums2[beforeSecond]
				beforeSecond++
			}
		} else if beforeSecond > lastSecond-1 {
			if numberCurrent <= nums1[beforeFirst] {
				numberCurrent = nums1[beforeFirst]
				beforeFirst++
			}
		}
	}

	if mulity > 0 {
		return float64(numberCurrent)
	}
	return (float64(numberCurrentPre) + float64(numberCurrent)) / 2
}
