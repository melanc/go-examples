package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch2(t *testing.T) {
	var a []int = []int{1, 2, 2, 5, 9, 9, 9, 13, 14, 14, 17, 23, 28}
	start, end := binarySearch2(a, 9)
	if start >= 0 {
		fmt.Println(start, end)
	} else {
		fmt.Println("not found")
	}

}

func binarySearch2(nums []int, target int) (int, int) {
	start, end := -1, -1
	if len(nums) == 0 {
		return start, end
	}

	left := 0
	right := len(nums) - 1 // 注意
	// 判断边界值
	if target < nums[left] || target > nums[right] {
		return start, end
	}

	for left <= right { // 注意
		// 判断边界值
		if nums[left] == target {
			start = left
			break
		}
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		start = left
	}

	left = 0
	right = len(nums) - 1
	for left <= right {
		// 判断边界值
		if nums[right] == target {
			end = right
			break
		}
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && nums[right] == target {
		end = right
	}

	return start, end
}
