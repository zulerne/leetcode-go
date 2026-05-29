// https://leetcode.com/problems/parallel-courses/description/
package parallelcourses

func MinimumSemesters(n int, relations [][]int) int {
	var res int
	var finished int
	graph := make(map[int][]int, n)
	inDegree := make(map[int]int, n)

	for _, r := range relations {
		graph[r[0]] = append(graph[r[0]], r[1])
		inDegree[r[1]]++
	}

	var queue []int
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
			finished++
		}
	}

	for len(queue) > 0 {
		levelSize := len(queue)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			for _, ngh := range graph[node] {
				inDegree[ngh]--
				if inDegree[ngh] == 0 {
					queue = append(queue, ngh)
					finished++
				}
			}
		}

		res++
	}

	if finished != n {
		return -1
	}
	return res
}
