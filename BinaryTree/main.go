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

func depthFirstRecursive(root *Node) []int {   // Recursive way
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
	fmt.Println(MaxRootToLeafPathRecursive(&a))

}