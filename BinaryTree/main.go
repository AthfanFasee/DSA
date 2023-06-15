package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

type stack []*Node

// Appends an elemend to stack. When using it incode we have to actually re-assign it to stack like this, stack = stack.Push(&a)
func (s stack) Push(v *Node) stack {
	return append(s, v)
}

// Removes the last element from stack and returns the new stack(slice) and the node which just got poped
func (s stack) Pop() (stack, *Node) { // in here stack is a type itself right? our custom type. it's not a variable or smthng
	l := len(s)
	return s[:l-1], s[l-1]
}

// Depth first traversel (move deeper into tree before u go vertically). For an ex: always go to d or e from b rather than c at first
// Normally a stack is used to solve thse problems

// Should return an array containing values of the tree in depth-first order
func depthFirst(root *Node) []int { // Normal way
	values := []int{}

	if root == nil {
		return values
	}

	stack := stack{root}

	for len(stack) > 0 {
		s, current := stack.Pop()
		stack = s
		values = append(values, current.Val)

		if current.Right != nil {
			stack = stack.Push(current.Right)
		}
		if current.Left != nil {
			stack = stack.Push(current.Left)
		}
	}

	return values
}

func depthFirstRecursive(root *Node) []int { // Recursive way
	if root == nil {
		return []int{}
	}

	right := depthFirstRecursive(root.Right)
	left := depthFirstRecursive(root.Left)

	result := []int{root.Val}
	result = append(result, left...)
	result = append(result, right...)

	return result
}

// Breadth first traversel (in a tree travel across rather than going deeper) For an ex: always go to c from b rather then d or e
// Nomrally we use a queue  for these problems

// Should return an array containing values of the tree in breadth-first order
func breadthFirst(root *Node) []int {
	values := []int{}

	if root == nil {
		return values
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		values = append(values, current.Val)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}

	}

	return values
}

// Find if a node exists in a binary tree
func find(root *Node, target *Node) bool {
	if root == nil {
		return false
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == target {
			return true
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return false
}

// Find if a node exists in a binary tree (Recursive)
func findRecursive(root *Node, target *Node) bool {
	if root == nil {
		return false
	}

	if root == target {
		return true
	}

	return findRecursive(root.Left, target) || findRecursive(root.Right, target)

}

// Find sum of all nodes in a binary tree
func Sum(root *Node) int {
	if root == nil {
		return 0
	}

	queue := []*Node{root}
	total := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		total += current.Val

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return total
}

// Find sum of all nodes in a binary tree (Recursive)
func sumRecursive(root *Node) int {
	if root == nil {
		return 0
	}
	return root.Val + sumRecursive(root.Left) + sumRecursive(root.Right)
}

// Find the lowest value in a binary tree (consider tree is non-empty)
func Min(root *Node) int {
	minValue := math.MaxInt
	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Val < minValue {
			minValue = current.Val
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
	}

	return minValue
}

func MinRecursive(root *Node) int {
	if root == nil {
		return math.MaxInt
	}

	return MinValue(root.Val, MinValue(MinRecursive(root.Left), MinRecursive(root.Right)))
}

func MinValue(int1, int2 int) int {
	if int1 < int2 {
		return int1
	} else {
		return int2
	}
}
func MaxValue(int1, int2 int) int {
	if int1 > int2 {
		return int1
	} else {
		return int2
	}
}

// Calculate sum of values in all root to leafe paths and return the max value among them
// First think what will I do if only a single node is given? I'll return it's value. Well that's a basecase
// Second think of a c node which got only left child. Now when we call line 233, when it's right child which is nill is passed, thats another base case
// Then think of a short tree as we covered all other edge cases, like a tree with 3 nodes. I want to see who's bigger in b and c and I want to add the bigger one to myself(a)
func MaxRootToLeafPathRecursive(root *Node) int {
	if root == nil {
		return math.MinInt
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return root.Val + MaxValue(MaxRootToLeafPathRecursive(root.Left), MaxRootToLeafPathRecursive(root.Right))
}

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node
func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + MaxValue(MaxDepth(root.Left), MaxDepth(root.Right))
}

// Given the root of a binary tree, return its maximum depth.(104)
// A binary tree's maximum depth is the number of nodes (not number of paths) along the longest path from the root node down to the farthest leaf node
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
}

func Max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// Find if both trees match
func SameTree(root1, root2 *Node) bool {
	if root1 == nil && root2 == nil { // this base case comes for nil children of leaf, or empty tree
		return true
	}
	if root1 == nil || root2 == nil { // Think of a base case when one root is nil
		return false
	}
	// Instead of both if statments above, we can also write this smart code

	// if root1 == nil || root2 == nil {  // When any of them is nil, just compare both
	//     return root1 == root2		// If incase both  r nil comparing them will return true
	// }								// If incase only one is nil, comparing them will return false

	if root1.Val != root2.Val { // when both aren't nil, compare values
		return false
	}

	return SameTree(root1.Left, root2.Left) && SameTree(root1.Right, root2.Right)
}

// Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot
// Leetcode 572, consider root and subroot aren't nill
func SubTree(root, subRoot *Node) bool {
	// These 2 checks r important if nodes can be nill

	// if subRoot == nil {
	//     return true
	// }
	// if root == nil {
	//     return false
	// }

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Val == subRoot.Val {
			if SameTree(current, subRoot) {
				return true
			}
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return false
}

// Gven the root of a binary tree, invert the tree, and return its root (226)
func invertTree(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		current.Right, current.Left = current.Left, current.Right

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
	}

	return root
}

func invertTreeRecursive(root *Node) *Node {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left

	invertTreeRecursive(root.Left)  // In here when a null node of maybe let's say "C" comes, it will return root on the first check
	invertTreeRecursive(root.Right) // the only confusing part is I'm not doing anything with the nil (root) it returns, unlike other functions
	// Bcs we only want to do the swapping in this problem
	return root // And all we want to return in the end is root itself which was given to us right.
}

func main() {
	a := Node{Val: 1}
	b := Node{Val: 2}
	c := Node{Val: 3}
	d := Node{Val: 4}
	e := Node{Val: 5}
	f := Node{Val: 6}

	// o := Node{Val: 9}

	a.Left = &b
	a.Right = &c
	b.Left = &d
	b.Right = &e
	c.Right = &f

	//           1
	//         /   \
	//       2      3
	//     /   \      \
	//   4       5      6

	// fmt.Println(depthFirst(&a))
	// fmt.Println(depthFirstRecursive(&a))
	// fmt.Println(breadthFirst(&a))
	// fmt.Println(find(&a, &d))
	// fmt.Println(findRecursive(&a, &o))
	// fmt.Println(Sum(&a))
	// fmt.Println(sumRecursive(&a))
	// fmt.Println(sumRecursive(&a))
	// fmt.Println(Min(&a))
	// fmt.Println(MinRecursive(&a))
	// fmt.Println(MaxRootToLeafPathRecursive(&a))
	fmt.Println(MaxDepth(&a))

}
