package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var a []int = []int{1, 2, 5, 9, 13, 14, 17, 23, 28}
	idx := binarySearch(a, 5)
	if idx >= 0 {
		fmt.Println(a[idx])
	} else {
		fmt.Println("not found")
	}

}

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1 // 注意

	for left <= right { // 注意
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
