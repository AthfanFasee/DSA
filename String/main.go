package main

import (
	"fmt"
	"strings"
)

// 242
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
// Basically they both gotta have exacly same characters, same amount
// Constraints:
// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.
func anagram(s, t string) bool {
	//This problem can be solved by sorting the characters, but sorting is computationally expensive.
	// We want a solution that loops once over each string. We are going to do the following: keep a slice of integers representing the amount of appearances of each of the 26 lowercase English letters:
	chars := make([]int, 26)

	// We loop once over s, incrementing the relevant number. We want chars[0] to refer to 'a', chars[1] to refer to 'b', etc.
	for _, v := range s {
		i := v - 'a' // For an ex : if 'a' is 123, then 123 - 123 will be 0th index of array, then we increament it by one for everytime 'a' charcter or rune appeares
		chars[i]++   // if b for an ex is 124, then 124 - 123 will be 1st index of array
	}

	// Then we do the same over t. We can maintain another slice, but we also can just decrement towards 0 again
	for _, v := range t {
		i := v - 'a'
		chars[i]--
	}

	// Now, if chars contains only zeros, we return true, otherwise false. This operation is O(1).
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}

// Bonus: What if the inputs contain Unicode characters?
// we can use a map

// We need to make a map of runes, because bytes are not adequae for all unicode characters. The rest is easy

func isAnagram(s string, t string) bool {
	chars := make(map[rune]int)

	for _, v := range s {
		chars[v]++
	}

	for _, v := range t {
		chars[v]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type
// ( ( ) ) below algorithm works for this too
func validParentheses(s string) bool {
	// This is an easy way to return false at beginning
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	// This problems needs a stack
	stack := []rune{}

	for _, c := range s {
		// Instead of hashmap, this is exactly same and easy
		switch c {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			// If one of the '(' '[' '{' came before, stack cannot be empty at 2nd step
			// When for an ex in this string "( ( ) )",  ')' comes as c, it should be in stack if the element before it was '('.
			// Bcs if the ammend before it was '(' it's added to the stack in case 1 and default case is skipped, so it's not removed.
			// We are checking both of those conditions here
			if len(stack) == 0 || c != stack[len(stack)-1] { // So this line basically makes sure, ')' got it's matching curly '(' before it comes in the string
				return false
			}
			// Remove the top element from stack, bcs at the end len of stack should be 0 when parantheses matched. What if at the end one of '(' comes and string ends? then stack won't be empty when, forloop ends
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// Valid Palindrom
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Constraints
// 1 <= s.length <= 2 * 105
// s consists only of printable ASCII characters
func isPalindrom(s string) bool {
	// Use two pointers
	left := 0           // Very first element index wise
	right := len(s) - 1 // Very last element index wise

	// We have to move them until left surpass right
	for left < right {
		// This will move the left pointer incase the current left isnt alpha numeric
		for left < right && !isAlphaNumeric(rune(s[left])) { // int32 or rune is fine here bur rune is faster
			left++
		}

		for left < right && !isAlphaNumeric(rune(s[right])) {
			right--
		}
		// After making sure left and right r in alphanumeric we are checking equality here
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}

		// Then continue with the loop
		left++
		right--
	}

	return true
}

// Helper func for isPalindrom
func isAlphaNumeric(asciCode rune) bool {
	return ('a' <= asciCode && asciCode <= 'z') ||
		('A' <= asciCode && asciCode <= 'Z') ||
		('0' <= asciCode && asciCode <= '9')
}

// Given a string s, find the length of the longest substring without repeating characters.
// We are gfonna use a sliding window to solve this
// Note that len("aaa") is 3(number of bytes), but to get the first byte in go we should say ("aaa")[0]. So dont get confused by this.
// Also len of array // ar := []int{1, 2, 3} is 3, but ar[2] is 3, and arr[3] is out of bond
// Both array and string are same in this manner
func lengthOfLongestSubstring(s string) int {
	// Set two pointers to beginning
	// One pointer stays, one pointer moves and expand the window
	left, right, result := 0, 0, 0

	hash := make(map[byte]bool) // As they are all ASCII characters, byte will represent character completely (string contains only Single byte characters)

	// "abac"
	for right < len(s) {
		if hash[s[right]] != true {
			hash[s[right]] = true // Mark the byte or character as seen
			// Right pointer moved, so minus left from that and add 1 to get correct length of substring
			result = max(result, right-left+1) // Look like this, if given string is "a" on the first attempt when right and left r 0, the length of longest substring should be 1 right?
			right += 1                         // Move right pointer
		} else {
			hash[s[left]] = false // This simply means ignore the first ever "a" and move the window
			left += 1             // When duplicate found in "aba", we didn't change right's value, that means on the next iteration right will again check this same caharcter and consider it in new unique array or next window
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character.
// You can perform this operation at most k times.
// Return the length of the longest substring containing the same letter you can get after performing the above operations
func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res, maxf, left := 0, 0, 0

	// bbaaa  k = 1
	for right, _ := range s {
		count[s[right]]++ // s[right] will return a byte. Then update it's count in hash

		maxf = max(maxf, count[s[right]])

		// Move left pointer
		if (right-left+1)-maxf > k { // Current winow - maximum appeared charcter should be less than or equal to k for this to be valid. So if it's bigger move left pointer
			// This order matters. Before u update left pointer value, update hash
			count[s[left]]-- // When we decrement, we dont have to update max variable as reducing max variable wont have any effect on end result.
			left++           // Bcs at the end we want the ,maximum length, so we dont care if maxf reduces
		}

		res = max(res, right-left+1)
	}

	return res
}

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// 49
type CharCount [26]int

func groupAnagrams(strs []string) [][]string {
	hash := make(map[CharCount][]string)

	for _, word := range strs {
		charCount := CharCount{}
		for _, r := range word {
			charCount[r-'a']++
		}

		hash[charCount] = append(hash[charCount], word)
	}

	var result [][]string
	for _, value := range hash {
		result = append(result, value)
	}

	return result
}

// OR
func groupAnagramsMethod2(strs []string) [][]string {
	hash := make(map[string][]string)

	for _, word := range strs {
		var countword [26]int

		for _, r := range word {
			countword[r-'a']++
		}
		hashkey := fmt.Sprint(countword)
		hash[hashkey] = append(hash[hashkey], word)

	}

	res := [][]string{}

	for _, values := range hash {
		res = append(res, values)
	}

	return res
}

// 76. Minimum Window Substring
func minWindow(s string, t string) string {
	left, right := 0, 0
	countT := make(map[byte]int)
	window := make(map[byte]int)
	distinctCharacterCount := 0
	minSubstring := ""

	// fill in countT map first
	for index := range t {
		countT[t[index]]++
	}

	for right < len(s) {
		window[s[right]]++

		if countT[s[right]] == window[s[right]] {
			distinctCharacterCount++
		}

		// Keep popping element from left side of window until this condition is no longer met
		for distinctCharacterCount == len(countT) {
			if minSubstring == "" {
				minSubstring = s[left : right+1] // end is excluded, hence end + 1
			}
			if right-left+1 < len(minSubstring) { // right-left+1 gives windows size
				minSubstring = s[left : right+1]
			}

			window[s[left]]--
			if window[s[left]] < countT[s[left]] {
				distinctCharacterCount--
			}

			// move the window
			left++
		}
		right++
	}

	return minSubstring
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abac"))
}
