package main

import "fmt"

// 34. Find First and Last Position of Element in Sorted Array
// Run binarcy search twice to find first and last occurence of target
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	start := -1
	for left <= right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			start = mid
			right = mid - 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// If start occurence of target is not found, end cannot be found too. So avoid 2nd binary search
	// If If start occurence of target is not found, start variable will remain as -1
	if start == -1 {
		return []int{-1, -1}
	}

	left, right = 0, len(nums)-1
	end := -1
	for left <= right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			end = mid
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return []int{start, end}
}

// 153. Find Minimum in Rotated Sorted Array
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	res := nums[0]

	// If you are confused about how do we decide until left <= right this loop should run
	// but not until left < right, then you can put like a very base case and see it like this [2,1]
	// if we only put left < right. For this base case the loop will only run once. And in our first iteration
	// mid := (left + right) / 2
	//	res = min(res, nums[mid]) these lines will update res to be nums[0]

	for left <= right {
		// While looping if the area we choose using pointers ( that sub array) if its somehow
		// Ends up as sorted properly we can return then. But incase we ever got a lefter res, we still need to check that condition.
		if nums[left] < nums[right] {
			res = min(res, nums[left]) // We cant just say res = nums[left]. This will be wrong at situations when only right pointer moves at very first iteration. Use this array as example and see [4, 5, 1, 2, 3]
			break
		}

		mid := (left + right) / 2
		res = min(res, nums[mid])
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 33. Search in Rotated Sorted Array
func search(nums []int, target int) int {
	// Find the pivot.
	left, right := 0, len(nums)-1

	// [4,5,6,7,0,1,2], target = 0
	// [3,4,0,1,2]
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	pivot := left

	// Regular binary search
	left, right = pivot, pivot-1+len(nums) // 4, 10
	for left <= right {
		mid := (left + right) / 2
		midVal := nums[mid%len(nums)]

		if midVal > target {
			right = mid - 1
		} else if midVal < target {
			left = mid + 1
		} else {
			// The index mid % n is returned as the result, which accounts for the circular nature of the array and ensures the correct mapping to the original array index
			return mid % len(nums)
		}
	}

	return -1
}

// Second Method
func searchTwo(nums []int, target int) int {
	pivot := findPivot(nums)
	var left int
	var right int

	// If the pivot index is -1, then the array is sorted properly (either not rotated, or rotated in a way it ended up sorted)
	// Hence we can search the entire array
	// [4,5,6,7,0,1,2], target = 0
	if pivot == -1 {
		left = 0
		right = len(nums) - 1
		// If the target is less than the first element in the array,
		// then the target must be in the right half of the array
	} else if target < nums[0] {
		left = pivot
		right = len(nums) - 1
		// If the target is greater than the first element in the array,
		// then the target must be in the left half of the array
	} else {
		left = 0
		right = pivot - 1
	}

	// Binary search for the target
	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// Check if the target is at the left index
	if nums[left] == target {
		return left
	}
	return -1
}

// Find the pivot index if it exists, return -1 otherwise
// The first step is to find the pivot index, which is the index where the array is rotated
func findPivot(nums []int) int {
	// Since all elements in the array are unique,
	// if the array is rotated, the first element will be greater than the last element.
	// Hence we can do an early return if the first element is less than the last element.
	if nums[0] < nums[len(nums)-1] {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		// If the element at the mid index is greater than the nums[right],
		// then the pivot index must be to the right of the mid index
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// Find an element in infinity array (logn time). return its index
// Infinty array means array's length shouldnt be used
// Idea is to search chunks of array starting length 2 and double it
func Infinity(nums []int, target int) int {
	start := 0
	end := 1

	// if got doubt use a simple example
	for target > nums[end] {
		newStart := end + 1
		end = end + (end-start+1)*2
		start = newStart
	}

	return binarySearch(nums, target, start, end)
}

func binarySearch(nums []int, target, start, end int) int {
	for start <= end {
		mid := start + (end-start)/2

		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

// 852. Peak Index in a Mountain Array

func main() {
	nums := []int{1, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(Infinity(nums, 6))
}
