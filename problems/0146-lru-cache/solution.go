// https://leetcode.com/problems/lru-cache/description/
package lrucache

type LRUCache struct {
	Capacity int
	Nodes    map[int]*Node
	Head     *Node
	Tail     *Node
}

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	nodes := make(map[int]*Node, capacity)

	return LRUCache{
		Capacity: capacity,
		Nodes:    nodes,
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) addToHead(node *Node) {
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Nodes[key]

	if !ok {
		return -1
	}

	this.remove(node)
	this.addToHead(node)

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Nodes[key]

	if ok {
		node.Value = value
		this.remove(node)
		this.addToHead(node)
		return
	}

	if len(this.Nodes) >= this.Capacity {
		delete(this.Nodes, this.Tail.Prev.Key)
		this.remove(this.Tail.Prev)
	}

	node = &Node{Key: key, Value: value}
	this.Nodes[key] = node
	this.addToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
