package main

import (
	"fmt"
)

// Find nth element in fib sequence
// Steps
// Get base cases (n ==1 and 2 always returns 1)
// n - 1 added with n - 2 (last 2 numbers) with recursive
func fib(n int, memo []int) int {
	// Check Memoized array
	if memo[n] != 0 {
		return memo[n]
	}
	
	if n == 1 || n == 2 {
		result := 1
		memo[n] = result //Memoize
		return result
	}
	result:= fib(n-1, memo) + fib(n-2, memo)
	memo[n] = result  // Memoize
	return result
}

// Find nth element in fib sequence (Bottum_up)
func fibBottomUp(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	
	bottumUp := make([]int, n+1)
	bottumUp[1] = 1
	bottumUp[2] = 1

		// All of these are constant time 
	for i := 3; i <= n; i++{
		bottumUp[i] = bottumUp[i-1] + bottumUp[i-2]
	}
	
	return bottumUp[n]

}

// Find nth element in fib sequence (Space : O(1))
// Instead of memoising using O(N) space to calculate the fibonacci number for n, there are two key observations:
// Fibonacci number calculation must start from fib-1 and fib-2, all the way to fib-n
// At each round k, we only need fib-(k-1) and fib-(k-2) for calculation
// As such, following the same dynamic programming spirit, we only need to iteratively calculate the numbers, starting from 1, while maintaining the fibonacci numbers for previous two rounds.
func fibLessSpace(n int) int {
	if n < 2 {
	  return n
	}
  
	/**
	 * Use a "sliding window" to keep track of
	 * the two numbers before current,
	 * and iterate for n-1 rounds until reaching n.
	 * prev1 and prev2 are initialised as
	 * fib-1 and fib-0 respectively.
	 * TC: O(N)
	 * SC: O(1)
	 */
	num, prev1, prev2 := 0, 1, 0
	for i := 2; i <= n; i++ {
	  num = prev1 + prev2
	  prev1, prev2 = num, prev1
	}
  
	return num
  }

// Grid Traveler
// Steps
// Get base cases (0 means can't reach at all), (1 x 1 means 1 way to reach)
// Go right(m -1) and left (n-1) in recursive calls and add them
func gridTraveler(m, n int, memo map[string]int) int {
	memoKey := fmt.Sprintf("%d,%d", m, n)
	if val, ok := memo[memoKey]; ok {
        return val
	}
	if (m == 0 || n == 0) {
		return 0
	}
	if (m == 1 && n == 1) {
		return 1
	}

	result := gridTraveler(m-1, n, memo) + gridTraveler(m, n-1, memo)
	memo[memoKey] = result
	return result
}

// Find if target sum can be generated from values of an array. An element can be used multiple times. For an ex : 3 can make 300 at the end 
// Note than if a single element matches target sum it still rturns true
// Steps
// Get base cases
// tareget - val in recursive
// return the main fucn when a rcursive returns true as for this problem, we dont have to check each possbile recursive call deep down
func canSum(targetSum int, values []int, memo map[int]bool) bool {
	if v, ok := memo[targetSum]; ok {
		return v
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, v := range values {
		reminder := targetSum - v
		if canSum(reminder, values, memo) {
			memo[targetSum] = true	
			return true
		}

	}
	memo[targetSum] = false
	return false
}

func main() {
	// memo := make([]int, 7)
	// fmt.Println(fib(6, memo))
	// fmt.Println(fibBottomUp(3))

	// memo := make(map[string]int)
	// fmt.Println(gridTraveler(3, 3, memo))

	// memo := make(map[int]bool)
	// fmt.Println(canSum(3, []int{2,5}, memo))


}