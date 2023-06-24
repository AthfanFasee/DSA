package main

import (
	"fmt"
)

type node struct {
	Val  string
	Next *node
}
type nodeInt struct {
	Val  int
	Next *nodeInt
}

// Travesel and print Values (Normal way)
func print(head *nodeInt) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

// Travesel and print Values (Recursive way)
func printRecursive(head *node) {
	if head == nil {
		return
	}

	fmt.Println(head.Val)
	printRecursive(head.Next)
}

// Append Values to an array and return it (Normal)
func createArray(head *node) []string {
	current := head
	s := []string{}
	for current != nil {
		s = append(s, current.Val)
		current = current.Next
	}
	return s
}

// Append Values to an array and return it (Recursive)
func createArrayRecursive(head *node) []string {
	s := []string{}
	fillValues(&s, head) // In GO, if you are appending something to a slice, you must pass it by reference. It's bcs when appending the slice header may change(len of slice or smthng), as everything is passed as value in go, the change on copy of the slice header wont be reflected in original slice header.
	return s
}

// Recursive helper for createArrayRecursive function
func fillValues(s *[]string, head *node) {
	if head == nil {
		return
	}
	*s = append(*s, head.Val) // Here we are appending. But incase we do something like s[0] = "Something" slice can be passed normally. In here no changes in slice header itself, just we are accessing the underlying array
	fillValues(s, head.Next)
}

// Find Sum of all node Val (Normal)
func sum(head *nodeInt) int {
	total := 0
	current := head
	for current != nil {
		total += current.Val
		current = current.Next
	}
	return total
}

// Find Sum of all node Val (Recursive)
func sumRecursive(head *nodeInt) int {
	if head == nil {
		return 0
	}
	return head.Val + sumRecursive(head.Next)
}

// Find if target Value exists in linkedList (Normal)
func exists(head *nodeInt, target int) bool {
	current := head
	for current != nil {
		if current.Val == target {
			return true
		}
		current = current.Next
	}
	return false
}

// Find if target Value exists in linkedList (Recursive)
func existsRecursive(head *nodeInt, target int) bool {
	if head == nil {
		return false
	}

	if head.Val == target {
		return true
	}

	return existsRecursive(head.Next, target)
}

// Find Value of the node in the given index (Normal)
func find(head *nodeInt, index int) int {
	current := head
	currentIndex := 0

	for current != nil {
		if index == currentIndex {
			return current.Val
		}
		current = current.Next
		currentIndex++
	}

	return 0
}

// Find Value of the node in the given index (Recursive)
func findRecursive(head *nodeInt, index int) int {
	if head == nil {
		return 0
	}
	if index == 0 {
		return head.Val
	}
	// index--   // OR we can reduce it here and pass it in function call as well.
	return findRecursive(head.Next, index-1) // index is being reduced by one here and being passed to function
}

// Reverse linked list and return the reversed linked list (means reference to head of returned liked list) (Normal)
func reverse(head *nodeInt) *nodeInt {
	var prev *nodeInt
	for head != nil {
		head.Next, prev, head = prev, head, head.Next // This single line can replace all of the below code. Underthe hood this will temporarily create variables and swap them, so we don't need to worry abt loosing head.Next on the first swap
		// Next := head.Next
		// head.Next = prev
		// prev = head
		// head = Next
	}
	return prev
}

// Merge two linked lists and return the merged linked list
// 1 -> 2 -> 3 -> 4 and 5 -> 6 merged should give us  1 -> 5 -> 2 -> 6 -> 3 -> 4

// keep track of current 1 and 2
// start from current 1
// as tail is alrdy head 1, current1's starting point should be current1.Next
// loop until current 1 and 2 are null
// get a counter and use it to add back and forth
// make progress on current1 and 2 inside of conditions as we only want progress if it's added to tail
// make progress of counter
// make progress of tail as the Next iteration should add to tail.Next's Next
// if one list ends add all rest of other list to tail.Next

func merge(head1 *nodeInt, head2 *nodeInt) *nodeInt {
	head := head1        // Keep track of head1 as it's gonna end up being the head of final result and, it's value is being changes later on
	result := &nodeInt{} // We don't have to worry abt the fact this linkedList starts with an empty struct(with zero values inside as 0 and nil). Doing this won't affect our end result
	// The reason for creating adummy node like this is bcs incase the lists they give us is empty, we dont need to worry abt edge cases. so just use it
	counter := 0

	for head1 != nil && head2 != nil {
		if counter%2 == 0 {
			result.Next = head1
			head1 = head1.Next
		} else {
			result.Next = head2
			head2 = head2.Next
		}
		result = result.Next
		counter++
	}

	if head1 != nil {
		result.Next = head1
	}
	if head2 != nil {
		result.Next = head2
	}
	return head
}

// 21. Merge Two Sorted Lists
// Merge the two lists in a one sorted list. Return the head of the merged linked list (Leetcode 21)
func mergeSorted(head1 *nodeInt, head2 *nodeInt) *nodeInt {
	result := &nodeInt{}
	tail := result
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val { // Add the node with min Val, first to the tail
			tail.Next = head1
			head1 = head1.Next
		} else {
			tail.Next = head2
			head1 = head2.Next
		}
		tail = tail.Next
	}

	if head1 != nil {
		tail.Next = head1
	}
	if head2 != nil {
		tail.Next = head2
	}

	return result.Next
}

// Given head, the head of a linked list, determine if the linked list has a cycle in it (Using map) (Leetcode 141)
func hasCycle(head *nodeInt) bool {
	nodesMap := make(map[*nodeInt]bool) // Use a node's pointer itself as key to map

	if head == nil { // if very first head is nil means, it cannot have a cycle at all. If we assume linked list cannot be empty, this line is not needed
		return false
	}

	for head != nil {

		if nodesMap[head] { // Check if the exact same pointer alrdy exists in map and return true if yes
			return true
		}

		nodesMap[head] = true // Otherwise add the pointer with Value true in map

		head = head.Next
	}

	return false
}

// Given head, the head of a linked list, determine if the linked list has a cycle in it (Using slow fast method)
func hasCycleSlowFast(head *nodeInt) bool {

	if head == nil { // if very first head is nil means, it cannot have a cycle at all. If we assume linked list cannot be empty, this line is not needed
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil { // Just try and loop until end

		slow = slow.Next      // slow jumps one pointer ahead
		fast = fast.Next.Next // fast jumps 2 pointer ahead. This is bcs in this way unless there's a cycle these 2 will never meet again OR fast will reach to null pretty soon

		if slow == fast { // If there's a cycle they will meet again for sure
			return true
		}
	}

	return false
}

// Given the head of a singly linked list, return the middle node of the linked list (link not empty) // 876
func findMidNode(head *nodeInt) *nodeInt {
	slow, fast := head, head
	for fast != nil && fast.Next != nil { // By the time fast will reach the end, slow will be at the middle of the linked List. Bcs the fast one goes twice as fast.
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect (the node they share in common). If the two linked lists have no intersection at all, return null
// If we know the distance of both lists, u can minus them and start the longer one with the diffrence ahead and travesal them and compare each node
// Instead of calculating their lengths,
// But a more optimal solution is, we just travesal both, when the shorter one reaches nul, we change the variable to point to longer one's head and iterate there
// Meanwhile when the longer one finally finishes it's iteration and it will start travesing shorter one, the current longer travesal will meet it, then just keep comparing and when they both match the same node return it.
// The longer one at it's second iteration will skip the difference in lenght (as it moves to short list), and the short one in it's 2nd iteration will iterate throught more nodes (more nodes will be equal to the difference in length as well as it moves to longer one this time)
// In this way we are somehow neglecting the difference in length (+2  -2) and they both will start looping as same length (like same lenght in nodes from the pic in leetcode) on 2nd iteration. (either they will meeet same node first or they will both become null first)
// When the 2nd iteration is done, if they both don't intersect, they both will reach null at the same time. So we can return nul (both b and a are nul at the moment so just return one of them)
func intersect(head1 *nodeInt, head2 *nodeInt) *nodeInt {
	a := head1
	b := head2
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = head2
		}
		if b != nil {
			b = b.Next
		} else {
			b = head1
		}
	}
	return a // When they are both equal or when they both are nil, it doesnt matter which one u return as answer
}

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well
// As this is sorted, just traversal through the list, and check each node with it's Next one.
// If match connect the current nodes Next to current.Next.Next or set it as usual to current.Next as we normally traversal
// As in sorted list duplicates can only appear Next to each other, it will work just fine
func deleteDuplicateSorted(head *nodeInt) *nodeInt {
	if head == nil {
		return nil
	} // Handle nil case
	current1 := head

	for current1 != nil && current1.Next != nil { // IF we don't check for current.Next != nil and try to access it's field like current.Next.Val, function will panic
		if current1.Val == current1.Next.Val {
			current1.Next = current1.Next.Next // This is how u delete the nodein between, just change the connection
		} else {
			current1 = current1.Next
		}
	}

	return head
}

// 143. Reorder List
func reorderList(head *node) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversed := reverse2(slow.Next)
	slow.Next = nil

	l1, l2 := head, reversed

	for l1 != nil && l2 != nil {
		nextL1, nextL2 := l1.Next, l2.Next
		l1.Next = l2
		l2.Next = nextL1
		l1, l2 = nextL1, nextL2
	}
}

func reverse2(head *node) *node {
	var previous *node
	for head != nil {
		head.Next, previous, head = previous, head, head.Next
	}

	return previous
}

// 19. Remove Nth Node From End of List
func removeNthFromEnd(head *node, n int) *node {
	// Create a new instance of the node struct and returns it's pointer
	dummy := new(node)
	dummy.Next = head
	left := dummy
	right := head

	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next
	return dummy.Next
}

// 23. Merge k Sorted Lists
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
func mergeKLists(lists []*node) *node {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		// pop 2 lists
		l1 := lists[0]
		l2 := lists[1]
		lists = lists[2:]

		merged := mergeTwoSortedListsRecursive(l1, l2)
		// The merged list is then appended to the end of lists slice
		lists = append(lists, merged)
	}

	return lists[0]
}

// This is also valid answer for leetcode 21 (merge two sorted linked lists)
func mergeTwoSortedListsRecursive(l1, l2 *node) *node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoSortedListsRecursive(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoSortedListsRecursive(l2.Next, l1)
		return l2
	}
}

func main() {
	// a := node{Val: "A"}
	// b := node{Val: "B"}
	// c := node{Val: "C"}
	// d := node{Val: "D"}

	// a.Next = &b
	// b.Next = &c
	// c.Next = &d

	a := nodeInt{Val: 1}
	b := nodeInt{Val: 2}
	c := nodeInt{Val: 2}
	d := nodeInt{Val: 4}

	a.Next = &b
	b.Next = &c
	c.Next = &d

	x := nodeInt{Val: 5}
	v := nodeInt{Val: 6}

	x.Next = &v
	// fmt.Println(test(&a))
	// print(&a)
	// printRecursive(&a)
	// fmt.Println(createArray(&a))
	// fmt.Println(createArrayRecursive(&a))
	// fmt.Println(sum(&a))
	// fmt.Println(sumRecursive(&a))
	// fmt.Println(exists(&a, 9))
	// fmt.Println(existsRecursive(&a, 4))
	// fmt.Println(find(&a, 1))
	// fmt.Println(findRecursive(&a, 5))
	// fmt.Println(print2(&a))
	// fmt.Println(reverse(&a))

	new := (merge(&a, &x))
	fmt.Println(new)

	// new := (mergeSorted(&a, &x))
	// fmt.Println(new)

	// fmt.Println(hasCycle(&a))
	// fmt.Println(findMidNode(&a))
	// fmt.Println(intersect(&a, &x))

	// new := (deleteDuplicateSorted(&a))
	// fmt.Println(new)
	// print(new)
}
