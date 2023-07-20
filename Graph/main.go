package main

import (
	"fmt"
	"math"
)

type stack []string

type graph map[string][]string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

// Print graph in depth-first order
func depthFirst2(graph map[string][]string, source string) {

}

func depthFirst(graph map[string][]string, source string) {
	stack := stack{source}

	for len(stack) > 0 {
		s, current := stack.Pop()
		stack = s
		fmt.Println(current)

		for _, v := range graph[current] {
			stack = stack.Push(v)
		}
	}
}

func depthFirstRecursive(graph map[string][]string, source string) {
	// A base case like this is not needed here, In Go, if you attempt to range over a nil slice, it won't result in a runtime error. Instead, the loop just won't execute, and it will automatically not call my recursive function at that scenario
	// if graph[source] == nil {
	// 	return
	// }

	fmt.Println(source)

	for _, value := range graph[source] {
		depthFirstRecursive(graph, value)
	}
}

// Print graph in breadth-first order
func breadthFisrt(graph map[string][]string, source string) {
	queue := []string{source}

	for len(queue) > 0 {
		current := queue[0]
		fmt.Println(current)
		queue = queue[1:]

		for _, v := range graph[current] {
			queue = append(queue, v)
		}
	}
}

// Find if we can travel from src to dst in graph
func hasPath(graph map[string][]string, src string, dst string) bool {
	queue := []string{src}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == dst {
			return true
		}

		for _, v := range graph[current] {
			queue = append(queue, v)
		}
	}

	return false
}

// Find if we can travel from src to dst in graph (recursive)
func hasPathRecsursive(graph map[string][]string, src string, dst string) bool {
	if src == dst {
		return true
	}

	for _, v := range graph[src] {
		if hasPathRecsursive(graph, v, dst) {
			return true
		}
	}

	return false
}

// Undirected Graph hasPath
// Find if these following edges can go from src to dst
func undirectedPathHasPath(edges [][]string, src string, dst string) bool {
	// Build graph () from given edges (Adjacency list using the edges given)
	graph := make(graph)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1]) // If edge[0] is not a key in graph, it will create it.
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return hasPathWithSet(graph, src, dst, make(map[string]bool))
}

// Helper recursive function for undirectedPath hasPath
func hasPathWithSet(graph graph, src string, dst string, visitedSet map[string]bool) bool {
	if src == dst {
		return true
	}

	if visitedSet[src] {
		return false
	}

	visitedSet[src] = true

	for _, v := range graph[src] {
		if hasPathWithSet(graph, v, dst, visitedSet) {
			return true
		}
	}

	return false
}

// Find count of connected components in graph
func connectedComponentCount(graph graph) int {
	count := 0
	visitedSet := make(map[string]bool)

	for Node := range graph {
		if explore(graph, Node, visitedSet) {
			count++
		}
	}

	return count
}

// Helper travesel function for connectedComponentCount
func explore(graph graph, Node string, visitedSet map[string]bool) bool {
	if visitedSet[Node] {
		return false
	}

	visitedSet[Node] = true

	for _, v := range graph[Node] {
		explore(graph, v, visitedSet)
	}

	return true
}

// Return the count of largest component in an undirected graph
func largestComponent(graph graph) int {
	longest := 0
	visitedSet := make(map[string]bool)

	for node := range graph {
		size := exploreWithCount(graph, node, visitedSet)
		if size > longest {
			longest = size
		}
	}

	return longest
}

func exploreWithCount(graph graph, node string, visitedSet map[string]bool) int {
	if visitedSet[node] {
		return 0
	}

	visitedSet[node] = true

	size := 1

	for _, v := range graph[node] {
		size += exploreWithCount(graph, v, visitedSet) // We shouldn't return here. Bcs in scenarios where this function is meant to explore the whole thing and find something (in here total count of nodes)
		// We shouldn,t be returning in middle. Dont consider the return will wait for recursion to end. For some reason it's messed up. But returning at last is fine
	}

	return size
}

// Find shortest path from NodeA to NodeB (edge count) in undirected graph.
func funcshortestPath(edges [][]string, nodeA string, nodeB string) int {
	graph := make(graph)
	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	vistedSet := make(map[string]bool)
	vistedSet[nodeA] = true

	queue := [][]interface{}{{nodeA, 0}}

	for len(queue) > 0 {
		current := queue[0][0].(string)
		distance := queue[0][1].(int)

		if current == nodeB {
			return distance
		}

		queue = queue[1:]

		for _, v := range graph[current] {
			if vistedSet[v] == false {
				vistedSet[v] = true
				queue = append(queue, []interface{}{v, distance + 1})
			}
		}
	}

	return -1
}

// Count "L" or lands or islands in given grid.
func islandCount(grid [][]string) int {
	count := 0

	for r, row := range grid {
		for c := range row {
			if exploreGrid(grid, r, c) {
				count++
			}
		}
	}

	return count
}

func exploreGrid(grid [][]string, r, c int) bool {
	rowInbounds := r >= 0 && r < len(grid)
	columnInbounds := c >= 0 && c < len(grid[0])
	if !rowInbounds || !columnInbounds {
		return false
	}

	if grid[r][c] == "W" {
		return false
	}

	// Instead of using visitedSet, you can simply change the original value to W as it's not needed after the above check
	grid[r][c] = "W"

	// According to the logic in islandCount and our basecase, only when a land comes, i'll start exploring using these recursives
	exploreGrid(grid, r-1, c)
	exploreGrid(grid, r+1, c)
	exploreGrid(grid, r, c+1)
	exploreGrid(grid, r, c-1)

	// So i can return true after a land is explored.
	return true
}

// Return length of smallest island in given grid.
func smallIslandSize(grid [][]string) int {
	count := math.MaxInt32 // Can be used over infinity
	for r, row := range grid {
		for c := range row {
			size := exploreIslandSize(grid, r, c)
			if size != 0 && size < count {
				count = size
			}
		}
	}

	return count
}

func exploreIslandSize(grid [][]string, r, c int) int {
	rowInbounds := r >= 0 && r < len(grid)
	columnInbounds := c >= 0 && c < len(grid[0])

	if !rowInbounds || !columnInbounds {
		return 0
	}

	if grid[r][c] == "W" {
		return 0
	}

	grid[r][c] = "W"

	size := 1

	size += exploreIslandSize(grid, r-1, c)
	size += exploreIslandSize(grid, r+1, c)
	size += exploreIslandSize(grid, r, c+1)
	size += exploreIslandSize(grid, r, c-1)

	return size
}

// 79. Word Search (also considered as backtracking problem)
// READ these notes before accessing this problem
// For ASCII characters, their rune and byte values are the same. This is because ASCII characters use only 7 bits,
// so they can be represented in a single byte, which has 8 bits.

// Rune in Go is a type that is used to represent a Unicode CodePoint.
// It can handle any Unicode character, which requires up to 4 bytes, not just ASCII characters.

// But when an ASCII character is represented as a rune, it still fits in a single byte,
// and the value is the same as the byte representation of the ASCII character.

func exist(board [][]byte, word string) bool {
	ROWS, COLUMNS := len(board), len(board[0])

	var dfs func(int, int, int) bool
	dfs = func(r, c, i int) bool {
		if len(word) == i {
			return true
		}

		if r < 0 || c < 0 || r >= ROWS || c >= COLUMNS {
			return false
		}
		if board[r][c] != word[i] || board[r][c] == '*' {
			return false
		}
		// It's same trick as explore island, but in here we only want a temporary change
		originalValue := board[r][c]
		// Conver this to something non-string, so it can never match a byte which represents a string
		// Converting here is fine, bcs we are alrdy done checking everything related to board[r][c]
		// In this case, the '*' character is an ASCII character, which can be represented as a single byte.
		// Thus, assigning this character to a byte is perfectly valid.
		// in here, Go automatically converts the rune to its byte representation because board[r][c] is of type byte.
		board[r][c] = '*'

		result := dfs(r+1, c, i+1) || dfs(r, c+1, i+1) || dfs(r-1, c, i+1) || dfs(r, c-1, i+1)

		// Changing back is important, bcs when we start from each place in grid for each chaarcter in word (word[i]),
		// we want it to be original again.
		board[r][c] = originalValue

		return result
	}

	for r, row := range board {
		for c := range row {
			if dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	// Directed Graph

	graph := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}
	// graph := map[int][]int{
	// 	3: {},
	// 	4: {6},
	// 	6: {4, 5, 7, 8},
	// 	8: {6},
	// 	7: {6},
	// 	5: {6},
	// 	1: {2},
	// 	2: {1},
	// }

	// Undirected Graph's edges

	// edges := [][]string{
	// 	{"i", "j"},
	// 	{"k", "i"},
	// 	{"m", "k"},
	// 	{"k", "l"},
	// 	{"o", "n"},
	// }

	// Adjacency list using the edges given (after we build graph from given edges)

	// graph := map[string][]string{
	// 	"i": {"j", "k"},
	// 	"j": {"i"},
	// 	"k": {"i", "m", "l"},
	// 	"m": {"k"},
	// 	"l": {"k"},
	// 	"o": {"n"},
	// 	"n": {"o"},
	// }

	// grid := [][]string{
	// 	{"W", "L", "W", "W", "W"},
	// 	{"W", "L", "W", "W", "W"},
	// 	{"W", "W", "W", "L", "W"},
	// 	{"W", "W", "L", "L", "W"},
	// 	{"L", "W", "W", "L", "L"},
	// 	{"L", "L", "W", "W", "W"},
	// }

	// depthFirst(graph, "a")
	depthFirst2(graph, "a")
	depthFirstRecursive(graph, "a")
	// breadthFisrt(graph, "a")
	// print(graph, "a")

	// fmt.Println(hasPath(graph, "a", "f"))
	// fmt.Println(hasPathRecsursive(graph, "a", "f"))

	// fmt.Println(undirectedPathHasPath(edges, "i", "o"))

	// fmt.Println(connectedComponentCount(graph))
	// fmt.Println(funcshortestPath(edges, "i", "o"))
	// fmt.Println(largestComponent(graph))
	// fmt.Println(funcshortestPath(edges, "i", "l"))
	// fmt.Println(islandCount(grid))
	// fmt.Println(smallIslandSize(grid))

}
