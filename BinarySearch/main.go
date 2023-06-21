package main

// 34. Find First and Last Position of Element in Sorted Array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target && (mid == len(nums)-1 || nums[mid+1] > target) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[left] != target {
		return []int{-1, -1}
	}

	rightBorderID := left

	left, right = 0, rightBorderID
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] >= target && (mid == 0 || nums[mid-1] <= target) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return []int{left, rightBorderID}
}
