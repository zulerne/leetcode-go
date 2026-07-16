// https://leetcode.com/problems/construct-quad-tree/description/
package constructquadtree

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var build func(row, col, size int) *Node
	build = func(row, col, size int) *Node {
		if size == 1 {
			return &Node{
				Val:    grid[row][col] == 1,
				IsLeaf: true,
			}
		}

		topLeft := build(row, col, size/2)
		topRight := build(row, col+size/2, size/2)
		bottomLeft := build(row+size/2, col, size/2)
		bottomRight := build(row+size/2, col+size/2, size/2)

		sameLeaf := topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf
		sameValue := topLeft.Val == topRight.Val && topLeft.Val == bottomLeft.Val && topLeft.Val == bottomRight.Val
		if sameLeaf && sameValue {
			return &Node{
				Val:    topLeft.Val,
				IsLeaf: true,
			}
		}

		return &Node{
			TopLeft:     topLeft,
			TopRight:    topRight,
			BottomLeft:  bottomLeft,
			BottomRight: bottomRight,
		}
	}

	return build(0, 0, len(grid))
}
