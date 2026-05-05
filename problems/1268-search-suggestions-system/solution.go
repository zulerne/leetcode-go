// https://leetcode.com/problems/search-suggestions-system//description/
package searchsuggestionssystem

import (
	"slices"
)

type Node struct {
	children [26]*Node
	words    []string
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for i := range len(word) {
		letter := word[i]
		child := cur.children[letter-'a']
		if child == nil {
			child = &Node{}
			cur.children[letter-'a'] = child
		}
		cur.words = append(cur.words, word)
		cur = child
	}
	cur.words = append(cur.words, word)
}

func (this *Trie) Search(word string) []string {
	cur := this.root
	for i := range len(word) {
		letter := word[i]
		child := cur.children[letter-'a']
		if child == nil {
			return nil
		}
		cur = child
	}
	return cur.words
}

func suggestedProducts(products []string, searchWord string) [][]string {
	suggested := make([][]string, len(searchWord))
	trie := Constructor()

	slices.Sort(products)
	for _, p := range products {
		trie.Insert(p)
	}

	for i := range len(searchWord) {
		words := trie.Search(searchWord[:i+1])

		if len(words) > 3 {
			suggested[i] = words[:3]
		} else {
			suggested[i] = words
		}
	}

	return suggested
}
