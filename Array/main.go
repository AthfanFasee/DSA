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
    indexMap := make(map[int]int)
	// loop through array
    for cIndex, cValue := range nums {
		// Find needed value
        neededValue := target - cValue
		// Check if it exists in hashmap yes means return it's index with currentArrayelement Index
		
        if targetIndex, ok := indexMap[neededValue]; ok {
            return []int{targetIndex, cIndex}
        }
        // No means add curentValue as key and index as value to hashmap
        indexMap[cValue] = cIndex
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
		if price - lowPrice > maxSale {
			maxSale = price - lowPrice
		}
	}

	// return the max difference
	return maxSale
}

// Given an integer array nums, find the subarray which has the largest sum and return its sum
func subArray(nums []int) int {
	max, sum := nums[0], nums[0]

	for _, v := range nums[1:] {  // Instead of trying to start from second index, just loop through array[:1]
		if sum >= 0 {
			sum += v // If sum is greater than or equal to 0 just add the current element to it too (as this new array which we find the sum for should be condigious, we won't miss anything by doing this)
		} else {
			sum = v // If the sum before this is less than zero, replace it with current element. As the negative sum won't help at all in finding maxsum
		}

		if max < sum { // For an example if [5,4,-8,1,2]. When -8 comes and reduces sum to 1, the max won't be replaces. It will be holding 9 which can be created by subarray [5,4].
			max = sum // So at the end we can return the max sum we ever got at any point of time
		}
	}

	return max
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
			return []int{left+1, right+1} // Add 1 bcs the array actually start at index 1 (question specific)
		} else if sum > target {  // If we surpass target, to reduce sum amount we should move right pointer towards left as array is sorted
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

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			target := nums[i] + nums[left] + nums[right]

			if target == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left, right = left+1, right-1

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
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

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	// fmt.Println(twoSum(nums, 17))
	// fmt.Println(Best(nums))
	// fmt.Println(subArray(nums))
	fmt.Println(threeSum(nums))
}