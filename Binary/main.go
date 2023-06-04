package main

// What is binary shifting? ( right shifting)
// The right-shift operation (num >>= 1) is equivalent to dividing the value by 2, but it performs integer division, which means it discards the remainder.
// Therefore, the result of the right-shift operation will always be an integer.
// In the case of the binary representation 10101, performing num >>= 1 will result in the new value of num being 1010.
// This corresponds to the decimal value 10, as you correctly pointed out.
// The right-shift operation does divide the value by 2, but it does not retain the fractional part.
// Instead, it simply drops the least significant bit, resulting in an integer value.
// So, in the context of the code and the right-shift operation (num >>= 1),
// the result will always be an integer value obtained by dividing num by 2 and rounding down the result.
// In the case of the binary representation 10101, the right-shift operation will produce the decimal value 10, not 10.5, as it discards the remainder during the division

// Left shifting is mulplying by2
// let's consider the decimal number 5, which has the binary representation 101.
// If we left-shift 5 by 2 positions (5 << 2), we get 10100, which is the binary representation of the decimal number 20

// LSB means Least Significant Bit (least bit)
// MSB means Most Significant bit which is the last 1 (after that 1 there could be meaningless 0s)

// Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).
func HammingWeight(num uint32) int {
	var res uint32

	for num != 0 {
		res += num & 1
		num >>= 1 // Right binary shifting
	}

	return int(res)
}

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array
func missingNumber(nums []int) int {
	// 0, 1, 3  Just apply this code on this example, and u'll get the idea
	result := 0

	for i, x := range nums {
		result ^= x // exclusive Or or XOR operator. Used to add 2 binaryu numbers without considering carry (carry means when 1 + 1, we add 0 then carry 1 to left)
		result ^= i + 1
	}

	return result
}

func missingNumberSumMethod(nums []int) int {
	sum := 0 // Calculate given array sum
	for _, num := range nums {
		sum += num
	}
	length := len(nums)
	total := length * (length + 1) / 2 // The formula to calculate the sum of consecutive numbers from 1 to n is (n * (n + 1)) / 2
	return total - sum                 // this difference will be missing Number
}

// Given two integers a and b, return the sum of the two integers without using the operators + and -.
func getSum(a int, b int) int {
	var carry int
	for b != 0 {
		carry = (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

// Reverse bits of a given 32 bits unsigned integer.
// Find each bit of number if each bit is 0 or 1 by (n & 1). This will give 1 if bit is 1 or will give 0
// Now that we have answer to add it to result (which alrdy got 32 0s by default), add with an Or operator
// Or operator is important to make sure we don't wrongly modify other bits of res
// Left Shift: res is left-shifted by 1. This shifts all the existing bits in res to the left, making room for the next bit.
// Bitwise OR: The least significant bit of n (obtained using n & 1) is ORed with res. This operation sets the rightmost bit of res to the same value as the current least significant bit of n.
// Right Shift: n is right-shifted by 1. This discards the least significant bit of n and prepares it for the next iteration.
func reverseBitsSecondWay(n uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		// res is fixed to 32 bits, so on left shift we will lose 0s from 32nd place but they dont mean anything
		res = res << 1 // Trick is on very first iteration this left shift dooesnt do anything as all res contains is zero but then res is being updated in next line and on 2nd iteration the updated res is moved a bit to the left
		res = res | (n & 1)
		n = n >> 1
	}
	return res
}

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i
// DP Offset O(N) |O(N)
func countBits(n int) []int {
	res := make([]int, n+1)
	offset := 1
	for i := 1; i <= n; i++ {
		if offset*2 == i {
			offset = i
		}
		res[i] = 1 + res[i-offset]
	}
	return res
}
