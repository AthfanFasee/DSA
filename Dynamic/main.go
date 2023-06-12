package main

import (
	"fmt"
	"strings"
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

	if n <= 2 {
		memo[n] = 1 //Memoize
		return 1
	}
	result := fib(n-1, memo) + fib(n-2, memo)
	memo[n] = result // Memoize
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
	for i := 3; i <= n; i++ {
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
	// We swap the values of m and n if m is greater than n. This ensures that the smaller value is always assigned to m and the larger value is assigned to n.
	// By doing this, the memo keys for scenarios like m = 3, n = 2 and m = 2, n = 3 will be the same: 2,3
	if m > n {
		m, n = n, m // Swap values if m is greater than n (this wont affect our end results)
	}

	memoKey := fmt.Sprintf("%d,%d", m, n)
	if val, ok := memo[memoKey]; ok {
		return val
	}
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
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
// target - val in recursive
// return the main fucn when a rcursive returns true as for this problem, we dont have to check each possbile recursive call deep down
func canSum(targetSum int, values []int, memo map[int]bool) bool {
	if v, ok := memo[targetSum]; ok {
		return v
	}
	if targetSum == 0 { // Even if the very first target itself is 0, its OK to return true
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, v := range values {
		reminder := targetSum - v
		if canSum(reminder, values, memo) { // As soon as i find true, or target sum reaches 0 in any call at all, I'm returning true early
			return true // So the only time I need memoization at a false only scenario.
		} // We can mark that wehn taget some is changed to sepcific values, that whole recursion tree will return false

	}
	memo[targetSum] = false // When a recursion function, finds the remaining targetsum's (mostly this will be ome middle level target) recursive tree returns false, memoize it
	return false            // Then return false to the caller
}

// Find Array elements that can generate target sum. return the very first combination u find
// Other rules r same as canSum
func howSum(targetSum int, values []int, memo map[int][]int) []int {
	if value, ok := memo[targetSum]; ok {
		return value
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, value := range values {
		remainder := targetSum - value
		remainderResult := howSum(remainder, values, memo)
		if remainderResult != nil {
			remainderResult = append(remainderResult, value)
			return remainderResult
		}
	}

	memo[targetSum] = nil
	return nil
}

// Return the shortest combination of array elements that can generate targetSum
func bestSum(targetSum int, values []int, memo map[int][]int) []int {
	if value, ok := memo[targetSum]; ok {
		return value
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	var shortestCombination []int

	for _, value := range values {
		remainder := targetSum - value
		remainderResult := bestSum(remainder, values, memo)
		if remainderResult != nil {
			remainderResult = append(remainderResult, value)
			if shortestCombination == nil {
				shortestCombination = remainderResult
			}
			if len(remainderResult) < len(shortestCombination) {
				shortestCombination = remainderResult
			}
		}
	}
	memo[targetSum] = shortestCombination
	return shortestCombination
}

// Find if given string values can construct the target string
func canConstruct(target string, wordBank []string, memo map[string]bool) bool {
	if value, ok := memo[target]; ok {
		return value
	}
	if target == "" {
		return true
	}

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]
			remainderResult := canConstruct(suffix, wordBank, memo)
			if remainderResult {
				return true
			}
		}
	}

	memo[target] = false
	return false
}

// Return the number of ways that the target can be constructed by
// Concatenating elements of the wordbank array
func countConstract(target string, wordBank []string, memo map[string]int) int {
	if value, ok := memo[target]; ok {
		return value
	}
	if target == "" {
		return 1
	}

	total := 0

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]
			numWaysForRest := countConstract(suffix, wordBank, memo)
			total += numWaysForRest
		}
	}

	memo[target] = total
	return total
}

// Return all possible combination of the strings from wordbank which makes up target string , in a 2D array
func allConstruct(target string, wordBank []string, memo map[string][][]string) [][]string {
	if value, ok := memo[target]; ok {
		return value
	}
	//{{}}: This creates an empty slice of strings []string{} and encloses it within another set of curly braces.
	// The double curly braces {{}} indicate that we are creating a slice with a single element,
	// and that element is itself an empty slice of strings
	if target == "" {
		return [][]string{{}} // [[]]
	}

	var result [][]string // This will just return an empty slice // []

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]
			suffixWays := allConstruct(suffix, wordBank, memo)
			for _, way := range suffixWays {
				targetWays := append([]string{word}, way...) // Same as saying in way array append the word variable's value. But doing like this makes sure we add the word to the front, which gives us nice order
				result = append(result, targetWays)
			}
		}
	}

	memo[target] = result
	return result
}

// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top
func climbStairs(n int) int {
	one, two := 1, 1

	for i := 0; i < n-1; i++ {
		one, two = one+two, one
	}

	return one
}

// 198. House Robber
func rob(nums []int) int {
	// [1, 4, 2]
	rob1, rob2 := 0, 0
	for _, n := range nums {
		temp := max(n+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}

	return rob2
}

// The only extra element here is, slice's first and last element r connected now.
// Everything else is exactly same as rob
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	// memo := make([]int, 7)
	// fmt.Println(fib(6, memo ))
	// fmt.Println(fibBottomUp(3))

	// memo := make(map[string]int)
	// fmt.Println(gridTraveler(3, 3, memo))

	// memo := make(map[int]bool)
	// memo := make(map[int][]int)
	// memo := make(map[string]bool)
	// memo := make(map[string]int)
	// memo := make(map[string][][]string)
	// fmt.Println(canSum(300, []int{7,14}, memo))
	// fmt.Println(howSum(14, []int{7,14}, memo))
	// fmt.Println(bestSum(70, []int{7, 14}, memo))
	// fmt.Println(canConstruct("ab", []string{"ab", "abc", "cd", "def", "abcd"}, memo))
	// fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed", []string{"eee", "e", "ee", "eee", "eee"}, memo))
	// fmt.Println(countConstract("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, memo))
	// fmt.Println(countConstract("purple", []string{"purp", "p", "ur", "le", "purpl"}, memo))
	// fmt.Println(countConstract("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed=", []string{"eee", "e", "ee", "eee", "eee"}, memo))
	// fmt.Println(allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, memo))
	// fmt.Println(allConstruct("aaaaaaaaaaaaaaazs", []string{"a", "aa", "aaa", "a", "aaaaa"}, memo))
}
