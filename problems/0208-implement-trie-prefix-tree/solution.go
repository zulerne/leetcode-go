// https://leetcode.com/problems/implement-trie-prefix-tree/description/
package implementtrieprefixtree

type Node struct {
	children   [26]*Node
	isTerminal bool
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
	for i := range len(word) {
		letter := word[i]
		child := cur.children[letter-'a']
		if child != nil {
			cur = child
			continue
		}
		child = &Node{
			children: [26]*Node{},
		}
		cur.children[letter-'a'] = child
		cur = child
	}
	cur.isTerminal = true
}

func (this *Trie) Search(word string) bool {
	node, ok := this.traverse(word)
	return ok && node.isTerminal
}

func (this *Trie) StartsWith(prefix string) bool {
	_, ok := this.traverse(prefix)
	return ok
}

func (this *Trie) traverse(word string) (*Node, bool) {
	cur := this.root
	for i := range len(word) {
		letter := word[i]
		child := cur.children[letter-'a']
		if child == nil {
			return nil, false
		}
		cur = child
	}
	return cur, true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
