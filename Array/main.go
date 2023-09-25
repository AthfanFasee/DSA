package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 2 Sum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	// loop through array
	for nIndex, nValue := range nums {
		// Find needed value
		neededValue := target - nValue
		// Check if it exists in hashmap yes means return it's index with currentArrayelement Index

		if targetIndex, ok := hashmap[neededValue]; ok {
			return []int{targetIndex, nIndex}
		}
		// No means add curentValue as key and index as value to hashmap
		hashmap[nValue] = nIndex
	}
	return []int{}
}

// You are given an array prices where prices[i] is the price of a given stock on the ith day
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0
func Best(prices []int) int {
	// All We need is the difference (maxSale)
	// Keep track of the least left value (to minus from every price from array). Start with very first element
	left := prices[0]
	maxSale := 0
	for _, price := range prices {
		// If any price is found as lesser than current left pointer (the value in the pointer), update it
		// There's only one rule to my left pointer which is to be the lowest array element so far
		if price < left {
			left = price
		}

		// Minus leftprice from each current price and keep trackof maximum difference in a variable
		if price-left > maxSale {
			maxSale = price - left
		}
	}

	// return the max difference
	return maxSale
}

// Given an integer array nums, find the subarray which has the largest sum and return its sum
func subArray(nums []int) int {
	maxSum, curSum := nums[0], nums[0]

	for _, v := range nums[1:] { // Instead of trying to start from second index, just loop through array[1:]
		if curSum >= 0 {
			curSum += v // If sum is greater than or equal to 0 just add the current element to it too (as this new array which we find the sum for should be condigious, we won't miss anything by doing this)
		} else {
			curSum = v // If the sum before this is less than zero, replace it with current element. As the negative sum won't help at all in finding maxsum
		}

		if maxSum < curSum { // For an example if [5,4,-8,1,2]. When -8 comes and reduces sum to 1, the max won't be replaces. It will be holding 9 which can be created by subarray [5,4].
			maxSum = curSum // So at the end we can return the max sum we ever got at any point of time
		}
	}

	return maxSum
}

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2
// The tests are generated such that there is exactly one solution. You may not use the same element twice
func twoSumsTwo(nums []int, target int) []int {
	// This is sorted so we can use 2 pointers to find our target sum which is more efficient
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left + 1, right + 1} // Add 1 bcs the array actually start at index 1 (question specific)
		} else if sum > target { // If we surpass target, to reduce sum amount we should move right pointer towards left as array is sorted
			right--
		} else {
			left++
		}
	}

	return []int{} // Assuming we have exactly one solution, the answere will be returned alrdy. This line is just to satisfy compiler
}

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k,
// and j != k, and nums[i] + nums[j] + nums[k] == 0
func threeSum(nums []int) [][]int {
	var res [][]int

	// Sort the array to use two pointers
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ { // We are gonna always be checking 3 elements in total, hence len(array) - 2
		if i > 0 && nums[i] == nums[i-1] { // From second iteration check if previous element is same as current, then skip as result cant contain duplicate
			continue // When continue is encountered in a loop, the remaining statements in the loop body are skipped, and the loop proceeds with the next iteration
		}

		// Now starts our 2 sum two
		left, right := i+1, len(nums)-1

		for left < right {
			target := nums[i] + nums[left] + nums[right]

			if target == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left, right = left+1, right-1

				for left < right && nums[left] == nums[left-1] {
					left++ // Skip duplicate once again (2th element on result)
				}

				for left < right && nums[right] == nums[right+1] {
					right-- // Skip duplicate once again (3rd element on result)
				}
			} else if target > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}

func ProductofArrayExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	// [1, 2, 3, 4]
	prefix := 1

	for i := range nums {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}

// Given an integer array nums, find a subarray that has the largest product, and return the product.
func maxProduct(nums []int) int {
	min, max, res := 1, 1, nums[0] // 1 <= nums.length <= 2 * 10^4
	for _, n := range nums {
		min, max = Min(n, Min(min*n, max*n)), Max(n, Max(min*n, max*n))
		res = Max(res, max)
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Incase Max or Min had to take multiple arguments
// Call is like this `var cur_min = minOf(num, prev_min * num, prev_max * num)`
func minOf(vars ...int) int {
	min := vars[0]

	for _, val := range vars {
		if val < min {
			min = val
		}
	}

	return min
}

// 128. Longest Consecutive Sequence
func longestConsecutive(nums []int) int {
	// Construct a set out of the nums array.
	numsSet := make(map[int]bool)
	for _, n := range nums {
		numsSet[n] = true
	}

	// The answer is stored here.
	maxSequenceLen := 0

	// Iterate through the set.
	for n := range numsSet {
		// We check if n-1 is in the set. If it is, then n is not the beginning of a sequence
		// and we go to the next number immediately.
		if _, ok := numsSet[n-1]; !ok {
			// Otherwise, we increment n in a loop to see if the next consecutive value is stored in nums.
			seqLen := 1
			for {
				if _, ok = numsSet[n+seqLen]; ok {
					seqLen++
					continue
				}
				// When the sequence is over, see if we did better than before.
				maxSequenceLen = Max(seqLen, maxSequenceLen)
				break
			}
		}
	}

	return maxSequenceLen
}

// Encode and Decode Strings
func Encode(stringList []string) string {
	res := ""
	// [Neet, Code]
	for _, s := range stringList {
		res += strconv.Itoa(len(s)) + "#" + s
	}

	return res
}

func Decode(str string) []string {
	i := 0
	res := []string{}

	// "4#Neet4#Code"
	for i < len(str) {

		length, _ := strconv.Atoi(string(str[i]))

		res = append(res, str[i+2:i+2+length])

		i = i + 2 + length
	}

	return res
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
	fmt.Println(left, right)
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
	fmt.Println(pivot)
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

// 11. Container With Most Water
func maxArea(heights []int) int {
	res := 0
	left, right := 0, len(heights)-1

	for left < right {
		area := (right - left) * min(heights[left], heights[right])
		res = Max(res, area)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	// fmt.Println(twoSum(nums, 17))
	// fmt.Println(Best(nums))
	// fmt.Println(subArray(nums))
	// fmt.Println(threeSum(nums))
	fmt.Println(search(nums, 0))
	// fmt.Println(searchTwo(nums, 0))
}
