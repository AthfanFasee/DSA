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
		result ^= x // exclusive Or or XOR operator
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
func reverseBits(n uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		res = res << 1
		res = res | (n & 1)
		n = n >> 1
	}
	return res
}
