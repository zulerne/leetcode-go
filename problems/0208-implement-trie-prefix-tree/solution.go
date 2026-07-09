// https://leetcode.com/problems/implement-trie-prefix-tree/description/
package implementtrieprefixtree

type Node struct {
	Val      byte
	End      bool
	children [26]*Node
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			children: [26]*Node{},
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root

	for i := range word {
		key := word[i] - 'a'
		child := cur.children[key]
		if child == nil {
			child = &Node{Val: word[i], children: [26]*Node{}}
			cur.children[key] = child
		}
		cur = child
	}

	cur.End = true
}

func (this *Trie) traverse(word string) (*Node, bool) {
	cur := this.root
	for i := range word {
		child := cur.children[word[i]-'a']
		if child == nil {
			return nil, false
		}
		cur = child
	}

	return cur, true
}

func (this *Trie) Search(word string) bool {
	node, ok := this.traverse(word)
	return ok && node.End
}

func (this *Trie) StartsWith(prefix string) bool {
	_, ok := this.traverse(prefix)
	return ok
}
