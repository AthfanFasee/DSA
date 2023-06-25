package main

// 211. Design Add and Search Words Data Structure
// To be most proficient here, we are gonna use Trie or Prefix Tree data structure
type TrieNode struct {
	IsWord   bool
	Children [26]*TrieNode
}

type WordDictionary struct {
	Root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{Root: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.Root
	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		if curr.Children[idx] == nil {
			curr.Children[idx] = &TrieNode{}
		}
		curr = curr.Children[idx]
	}
	curr.IsWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return Search(0, word, this.Root)
}

func Search(idx int, word string, node *TrieNode) bool {
	// If idx equals len(word), which means we've processed all characters and reached the end of the word.
	// It checks IsWord for the current node and returns its value
	if idx == len(word) { // Array index starts at 0. So this check is idx == len(word) and not index == len(word) -1
		return node.IsWord
	}

	if word[idx] == '.' {
		for i := 0; i < 26; i++ { // If wildcard is found, loop through all possible children.
			if node.Children[i] != nil && Search(idx+1, word, node.Children[i]) { // Unless they are not nill call Search for all of them
				return true
			}
		}
	} else {
		i := int(word[idx] - 'a')
		if node.Children[i] == nil {
			return false
		}
		return Search(idx+1, word, node.Children[i])
	}

	return false
}
