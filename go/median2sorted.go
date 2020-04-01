// https://leetcode.com/problems/median-of-two-sorted-arrays/
package main

import (
	"fmt"
)

func mergeSorted(nums1 []int, nums2 []int) []int {
	i, i1, i2 := 0, 0, 0
	out := make([]int, len(nums1)+len(nums2))
	for ; i < len(nums1)+len(nums2); i++ {
		if i1 >= len(nums1) || i2 >= len(nums2) {
			break
		}
		if nums1[i1] < nums2[i2] {
			out[i] = nums1[i1]
			i1++
		} else {
			out[i] = nums2[i2]
			i2++
		}
	}

	if i1 < len(nums1) {
		copy(out[i:], nums1[i1:])
	}
	if i2 < len(nums2) {
		copy(out[i:], nums2[i2:])
	}
	return out
}

func median(nums []int) float64 {
	i := len(nums) / 2
	if len(nums)%2 == 0 {
		return (float64(nums[i-1]) + float64(nums[i])) / 2
	}
	return float64(nums[i])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := mergeSorted(nums1, nums2)
	return median(nums)
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
