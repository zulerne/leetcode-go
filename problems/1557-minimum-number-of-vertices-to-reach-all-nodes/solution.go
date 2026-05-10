// https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/description/
package minimumnumberofverticestoreachallnodes

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var res []int
	inDegree := make([]int, n)

	for _, edge := range edges {
		inDegree[edge[1]]++
	}

	for i, v := range inDegree {
		if v == 0 {
			res = append(res, i)
		}
	}

	return res
}
