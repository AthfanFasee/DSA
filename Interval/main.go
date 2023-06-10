package main

import (
	"fmt"
	"sort"
)

// Meeting Rooms - Leetcode 252
//Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...]
// (si < ei), determine if a person could attend all meetings.
func canAttendMeetings(intervals [][]int) bool {
	// Sort the intervals based on the start time
	// The sort.Slice function in Go is a convenient way to sort slices. It takes the slice to be sorted (intervals in this case) as the first argument. The second argument is a function that defines the comparison logic for sorting.
	// This function takes two indices i and j and returns a boolean value indicating whether intervals[i] should come before intervals[j] in the sorted order.
	// In this case, the comparison function func(i, j int) bool { return intervals[i][0] < intervals[j][0] } compares the start time of the intervals. It accesses the start time of the intervals at indices i and j using intervals[i][0] and intervals[j][0] respectively.
	// By returning true when the start time of intervals[i] is less than the start time of intervals[j], the function ensures that intervals with earlier start times appear before intervals with later start times in the sorted slice
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Check if there is any overlap between consecutive meetings
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] { // The start time of current one cant be less than the end time of previous one
			return false // Overlap found, cannot attend all meetings
		}
	}

	return true // No overlap, can attend all meetings
}

//  Meeting Rooms - Leetcode 252
func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	startTime := make([]int, n)
	endTime := make([]int, n)

	// Separate the start times and end times into separate arrays
	for i := 0; i < n; i++ {
		startTime[i] = intervals[i][0]
		endTime[i] = intervals[i][1]
	}

	// Sort the start times and end times in ascending order
	sort.Ints(startTime)
	sort.Ints(endTime)

	rooms := 0
	endIndex := 0

	// Iterate over the start times
	for i := 0; i < n; i++ {
		if startTime[i] < endTime[endIndex] {
			// A new room is needed as the current meeting overlaps with a previous one
			rooms++
		} else {
			// The current meeting can reuse an existing room
			endIndex++
		}
	}

	return rooms
}

func main() {
	fmt.Println(canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(canAttendMeetings([][]int{{5, 8}, {9, 15}, {1, 4}}))
}
