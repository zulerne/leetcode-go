// https://leetcode.com/problems/design-linked-list/description/
package designlinkedlist

import "fmt"

type MyLinkedList struct {
	head *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: &Node{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.size {
		return -1
	}

	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	cur := this.head
	for range index {
		cur = cur.next
	}
	cur.next = &Node{val: val, next: cur.next}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}

	cur := this.head
	for range index {
		cur = cur.next
	}
	cur.next = cur.next.next
	this.size--
}

func (this *MyLinkedList) Print() {
	result := make([]int, 0)
	cur := this.head.next
	for cur != nil {
		result = append(result, cur.val)
		cur = cur.next
	}
	fmt.Println(result)
}
