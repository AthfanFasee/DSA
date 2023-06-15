package main

import (
	"fmt"
	"sort"
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
	// Keep track of the lowest value (to minus from every price from array). Start with very first element
	lowPrice := prices[0]
	maxSale := 0
	for _, price := range prices {
		// If any price is found as lower than current lower, update it
		if price < lowPrice {
			lowPrice = price
		}

		// Minus lowprice from each current price and keep trackof maximum difference in a variable
		if price-lowPrice > maxSale {
			maxSale = price - lowPrice
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
			curSum = v // If the sum before this is lesst han zero, replace it with current element. As the negative sum won't help at all in finding maxsum
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

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	// fmt.Println(twoSum(nums, 17))
	// fmt.Println(Best(nums))
	// fmt.Println(subArray(nums))
	fmt.Println(threeSum(nums))
}
