// https://leetcode.com/problems/design-add-and-search-words-data-structure/description/
package designaddandsearchwordsdatastructure

type WordDictionary struct {
	root *Node
}

type Node struct {
	end      bool
	children [26]*Node
}

func Constructor() WordDictionary {
	wd := WordDictionary{
		root: &Node{},
	}

	return wd
}

func (wd *WordDictionary) AddWord(word string) {
	cur := wd.root

	for i := range word {
		child := cur.children[word[i]-'a']
		if child == nil {
			child = &Node{}
			cur.children[word[i]-'a'] = child
		}

		cur = child
	}

	cur.end = true
}

func (wd *WordDictionary) Search(word string) bool {
	return wd.root.search(word, 0)
}

func (n *Node) search(word string, pos int) bool {
	if n == nil {
		return false
	}
	if pos == len(word) {
		return n.end
	}

	if word[pos] == '.' {
		for _, c := range n.children {
			if c != nil {
				if c.search(word, pos+1) {
					return true
				}
			}
		}
		return false
	}

	return n.children[word[pos]-'a'].search(word, pos+1)
}
