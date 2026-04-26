// https://leetcode.com/problems/design-hashmap/description/
package designhashmap

type Node struct {
	Key  int
	Val  int
	Next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Put(key, val int) {
	cur := ll.head
	for cur != nil {
		if cur.Key == key {
			cur.Val = val
			return
		}
		cur = cur.Next
	}
	node := &Node{Key: key, Val: val, Next: ll.head}
	ll.head = node
}

func (ll *LinkedList) Get(key int) int {
	cur := ll.head
	for cur != nil {
		if cur.Key == key {
			return cur.Val
		}
		cur = cur.Next
	}
	return -1
}

func (ll *LinkedList) Remove(key int) {
	dummy := &Node{Next: ll.head}
	cur := dummy
	for cur != nil && cur.Next != nil {
		if cur.Next.Key == key {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	ll.head = dummy.Next
}

type MyHashMap struct {
	n       int
	buckets []*LinkedList
}

func Constructor() MyHashMap {
	n := 997
	buckets := make([]*LinkedList, n)
	for i := range n {
		buckets[i] = &LinkedList{}
	}
	return MyHashMap{
		n:       n,
		buckets: buckets,
	}
}

func (this *MyHashMap) hash(x int) int {
	return x % this.n
}

func (this *MyHashMap) Put(key int, val int) {
	n := this.hash(key)
	this.buckets[n].Put(key, val)
}

func (this *MyHashMap) Get(key int) int {
	n := this.hash(key)
	return this.buckets[n].Get(key)
}

func (this *MyHashMap) Remove(key int) {
	n := this.hash(key)
	this.buckets[n].Remove(key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
