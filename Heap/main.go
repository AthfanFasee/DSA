package main

// 347
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
// This can be solved using bucketshort more efficiently than using a heap
// So can consider this as a normal array problem as well
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	bucket := make([][]int, len(nums)+1)

	for number, count := range freq {
		bucket[count] = append(bucket[count], number)
	}

	res := []int{}
	for i := len(bucket) - 1; i > 0; i-- {
		if len(res) < k {
			res = append(res, bucket[i]...)
		}
	}

	return res
}
