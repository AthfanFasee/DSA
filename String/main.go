package main

import (
	"fmt"
	"strings"
)

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
// Basically they both gotta have exacly same characters, same amount
//Constraints:
// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.

func anagram(s, t string) bool {
	//This problem can be solved by sorting the characters, but sorting is computationally expensive.
	// We want a solution that loops once over each string. We are going to do the following: keep a slice of integers representing the amount of appearances of each of the 26 lowercase English letters:
	chars := make([]int, 26)
    
	// We loop once over s, incrementing the relevant number. We want chars[0] to refer to 'a', chars[1] to refer to 'b', etc.
    for _, v := range s {
		i := int(v - 'a') // For an ex : if 'a' is 123, then 123 - 123 will be 0th index of array, then we increament it by one for everytime 'a' charcter or rune appeares
		chars[i]++		// if b for an ex is 124, then 124 - 123 will be 1st index of array
	}
    
	// Then we do the same over t. We can maintain another slice, but we also can just decrement towards 0 again
    for _, v := range t {
		i := int(v - 'a')
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
// We could maybe go on a quest to find out how many Unicode chars are there, and maintain a slice with many thousands of elements, and technically this would still be O(n), 
// but it is a cheeky solution with a lot of waste. We can do better, we can use a map

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
			// When for an ex ')' comes as c, it should be in stack if the ement before it was '('
			// We are checking both of those conditions here
            if len(stack) == 0 || c != stack[len(stack)-1] {
                return false
            }
			// Remove the top element from stack, bcs at the end len of stack should be 0. What if at the end one of '(' comes and string ends? them stack won't be empty when for loop ends
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}

// Valid Palindrom
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, 
// it reads the same forward and backward. Alphanumeric characters include letters and numbers
// Constraints
// 1 <= s.length <= 2 * 105
// s consists only of printable ASCII characters
func isPalindrom(s string) bool {
    // Use two pointers
    left := 0
    right := len(s) - 1

    // We have to move them until left surpass right
    for left < right {
        // As this s is only ASCII, it;s byte's decimal value(uint8) direclty is equal to rune (or unicode) value too
        for left < right && !isAlphaNumeric(rune(s[left])) { // int32 or rune is fine here bur rune is faster
            left++
        }

        for left < right && !isAlphaNumeric(rune(s[right])) {
            right--
        }

        if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
            return false
        }
        
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
            result = max(result, right - left + 1) // Look like this, if given string is "a" on the first attempt when right and left r 0, the length of longest substring should be 1 right? 
            right += 1  // Move right pointer
        } else {
            hash[s[left]] = false // This simply means ignore the first ever "a" and move the window
            left += 1    // When duplicate found in "aba", we didn't change right's value, that means on the next iteration right will again check this same caharcter and consider it in new unique array or next window
        }
    }
    
    return result
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. 
// You can perform this operation at most k times.
// Return the length of the longest substring containing the same letter you can get after performing the above operations
func characterReplacement(s string, k int) int {
    count := make([]int, 128) // 128 bytes to represent 128 ASCII characters.
    res := 0
    maxf := 0
    left := 0

    // bbaaa  k = 1
    for right := 0; right < len(s); right++ {
        count[s[right]] += 1    // s[end] will return a byte, which is the same as that element in the count array

        maxf = max(maxf, count[s[right]])
        
        // Move left pointer
        if (right - left + 1) - maxf > k  { // Current winow - maximum appeared charcter should be less than k for this to be valid. So if it's bigger move left pointer
            count[s[left]] -= 1  // When we decrement, we dont have to update max variable as reducing max variable wont have any effect on end result.
            left += 1            // Bcs at the end we want the ,aximum length, so we dont care if maxf reduces 
        }

        res = max(res, right - left + 1)
    }

    return res
}

func main() {
   fmt.Println(lengthOfLongestSubstring("abac"))
}