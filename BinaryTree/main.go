package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
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
// First think what will I do if only a single node is given? I'll return it's value. Well that's a basecase (if root.Left == nil && root.Right == nil)
// Second think of a c node which got only left child. Now When it's right child which is nill is passed, thats another base case (if root == nil)
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

// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST
// All Node.val are unique
// p and q will exist in the BST.
// p != q
// A binary search tree is a type of binary tree that follows a specific property:
// for every node, all values in its left subtree are smaller than its value, and
// all values in its right subtree are greater than its value.
func lowestCommonAncestor(root, p, q *Node) *Node {
	for root != nil {
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return root
}

func lowestCommonAncestorRecursive(root, p, q *Node) *Node {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorRecursive(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorRecursive(root.Right, p, q)
	}

	return root
}

// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*Node{root}

	result := [][]int{}

	for len(queue) > 0 {
		len := len(queue)
		levelNodes := []int{} // The idea is to figure out when to make this levelNodes empty again

		for i := 0; i < len; i++ {
			curr := queue[0]
			queue = queue[1:]

			levelNodes = append(levelNodes, curr.Val)

			// The question is asking for left to right. So left should be append to queue first
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		result = append(result, levelNodes)
	}

	return result
}

// Binary Tree Maximum Path Sum (Leetcode 124)
func maxPathSum(root *Node) int {
	globalMax := math.MinInt
	dfs(root, &globalMax)
	return globalMax
}

func dfs(root *Node, globalMax *int) int {
	if root == nil {
		return 0
	}

	pathSumFromLeft := max(dfs(root.Left, globalMax), 0)
	pathSumFromRight := max(dfs(root.Right, globalMax), 0)

	*globalMax = max(*globalMax, root.Val+pathSumFromLeft+pathSumFromRight)

	return root.Val + max(pathSumFromLeft, pathSumFromRight)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 129. Sum Root to Leaf Numbers
// You are given the root of a binary tree containing digits from 0 to 9 only.
func sumNumbers(root *Node) int {
	return sumNodes(root, 0)
}

func sumNodes(node *Node, num int) int {
	if node == nil {
		return 0
	}

	num = num*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return num
	}

	return sumNodes(node.Left, num) + sumNodes(node.Right, num)
}

// 105. Construct Binary Tree from Preorder and Inorder Traversal
func buildTree(preorder []int, inorder []int) *Node {
	if len(preorder) == 0 {
		return nil
	}
	//preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	idx := indexOf(inorder, preorder[0])
	return &Node{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
}

func indexOf(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

// 98. Validate Binary Search Tree
func isValidBST(root *Node) bool {
	return valid(root, math.MinInt, math.MaxInt)
}

func valid(node *Node, left, right int) bool {
	if node == nil {
		return true
	}
	if node.Val <= left || node.Val >= right {
		return false
	}
	return valid(node.Left, left, node.Val) && valid(node.Right, node.Val, right)
}

// 230. Kth Smallest Element in a BST
func kthSmallest(root *Node, k int) int {
	n := 0
	stack := []*Node{}
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n++
		if n == k { // According to question we are gurrenteed to return here
			return cur.Val
		}

		cur = cur.Right
	}

	return 0 // Just to satisy compiler
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *Node) string {
	var buffer bytes.Buffer

	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			buffer.WriteString("N,")
		} else {
			buffer.WriteString(strconv.Itoa(node.Val))
			buffer.WriteString(",")
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)

	return buffer.String()
}

func (this *Codec) deserialize(data string) *Node {
	tokens := strings.Split(data, ",")

	var dfs func() *Node
	dfs = func() *Node {
		token := tokens[0]
		tokens = tokens[1:]
		if token == "N" {
			return nil
		}
		val, _ := strconv.Atoi(token)
		return &Node{val, dfs(), dfs()}
	}

	return dfs()
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
	fmt.Println(breadthFirst(nil))
	// fmt.Println(find(&a, &d))
	// fmt.Println(findRecursive(&a, &o))
	// fmt.Println(Sum(&a))
	// fmt.Println(sumRecursive(&a))
	// fmt.Println(sumRecursive(&a))
	// fmt.Println(Min(&a))
	// fmt.Println(MinRecursive(&a))
	// fmt.Println(MaxRootToLeafPathRecursive(&a))
	// fmt.Println(MaxDepth(&a))

}
