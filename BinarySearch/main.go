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
	// [3,0,1,2]
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// We can slo say this is right. Bcs at the end of loop, as condition is left < right
	// both left and right will end up pointing same lement. The next iteration loop will break right. So yh
	// Both left and right now points to same guy which is the pivot
	pivot := left

	// Regular binary search
	// [4,5,6,7,0,1,2] // We need to binary search from 0 to 7 (thinking like imaginary array)
	left, right = pivot, pivot-1+len(nums) // 4, 10
	for left <= right {
		mid := (left + right) / 2     // This will give mid in imaginary array
		midVal := nums[mid%len(nums)] // This is how u find mid correctly in original array from imaginary array

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
		mid := left + (right-left)/2

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

// find pivot in sorted array which contains duplicates
func findPivotDuplicate(nums []int) int {
	// Since all elements in the array are unique,
	// if the array is rotated, the first element will be greater than the last element.
	// Hence we can do an early return if the first element is less than the last element.
	if nums[0] < nums[len(nums)-1] {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		// If the element at the mid index is greater than the nums[right],
		// then the pivot index must be to the right of the mid index
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			// nums[mid] == nums[right] // Earlier we didnt consider this case. But this time it's possible
			// To handle this case, we make sure if right is pivot then we cautiously decrease right by 1.
			// This is a safe move to make because it reduces the search space while ensuring that the pivot remains within it
			// the pivot index is the position of the smallest element in the array and the only place in the sorted and rotated array where the next number is smaller than the previous one
			// The right > 0 check is necessary to avoid an ArrayIndexOutOfBounds when right is 0
			if right > 0 && nums[right] < nums[right-1] {
				return right
			}
			right--
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
func peakIndexInMountainArray(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right { // Condition will be like this when at the end we expect both left and right to point at the answer. i dont want loop to go further like left being equal to left = right + 1 and all
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			// you are in dec part of array
			// this may be the ans, but look at left
			// this is why end != mid - 1
			right = mid
		} else {
			// you are in asc part of array
			left = mid + 1 // because we know that mid+1 element > mid element
		}
	}
	// in the end, start == end and pointing to the largest number because of the 2 checks above
	// start and end are always trying to find max element in the above 2 checks
	// hence, when they are pointing to just one element, that is the max one because that is what the checks say
	// more elaboration: at every point of time for start and end, they have the best possible answer till that time
	// and if we are saying that only one item is remaining, hence cuz of above line that is the best possible answer.
	return left // or return right. They both will point to answer
}

// 1095. Find in Mountain Array
// Find peak, binary search to asc array and desc array seprately
func findInMountainArray(target int, nums []int) int {
	peak := peakIndexInMountainArray(nums)
	firstTry := orderAgnosticBS(nums, target, 0, peak)
	if firstTry != -1 {
		return firstTry
	}
	return orderAgnosticBS(nums, target, peak+1, len(nums)-1)
}

// BinarySearch but Array could be in either ASC or DCS Order
func orderAgnosticBS(nums []int, target, start, end int) int { // normally u wont provide start and end, i'm providing as this is used as a helper

	// find whether the array is sorted in ascending or descending
	isAsc := nums[start] < nums[end]

	for start <= end {
		// find the middle element
		// int mid = (start + end) / 2; // might be possible that (start + end) exceeds the range of int in a programming language
		mid := start + (end-start)/2

		if nums[mid] == target {
			return mid
		}

		if isAsc {
			if target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			// for dcs order just change arrow direction :)
			if target > nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(findPivotDuplicate([]int{3}))
}
