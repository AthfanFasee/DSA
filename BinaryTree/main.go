package main

import "fmt"

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

type stack []*Node

func (s stack) Push(v *Node) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, *Node) {
	l := len(s)
	return s[:l-1], s[l-1]
}

// Depth first traversel (move deeper into tree before u go vertically). For an ex: always go to d or e from b rather than c at first
// Normally a stack is used to solve thse problems

// Should return an array containing values of the tree in depth-first order
func depthFirst(root *Node) []int {
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


func main() {
	a := Node{Val: 1}
	b := Node{Val: 2}
	c := Node{Val: 3}
	d := Node{Val: 4}
	e := Node{Val: 5}
	f := Node{Val: 6}

	a.Left = &b
	a.Right = &c
	b.Left = &d
	b.Right = &e
	c.Right = &f

	//           a
	//         /   \
	//       b      c
	//     /   \      \
	//   d       e      f

	// fmt.Println(depthFirst(&a))
	// fmt.Println(breadthFirst(&a))
	// fmt.Println(find(&a, &d))
	// fmt.Println(findRecursive(&a, &o))
	// fmt.Println(Sum(&a))
	// fmt.Println(sumRecursive(&a))
	fmt.Println(sumRecursive(&a))

}