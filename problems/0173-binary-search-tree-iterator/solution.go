// https://leetcode.com/problems/binary-search-tree-iterator/description/
package binarysearchtreeiterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{
		stack: make([]*TreeNode, 0),
	}

	it.pushLeft(root)

	return it
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if node.Right != nil {
		this.pushLeft(node.Right)
	}

	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
