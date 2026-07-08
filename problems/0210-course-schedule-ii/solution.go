// https://leetcode.com/problems/course-schedule-ii/description/
package coursescheduleii

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0, numCourses)
	inDegree := make([]int, numCourses)
	graph := make([][]int, numCourses)

	for _, p := range prerequisites {
		inDegree[p[0]]++
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	queue := []int{}

	for node, val := range inDegree {
		if val == 0 {
			queue = append(queue, node)
			res = append(res, node)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, ngh := range graph[node] {
			inDegree[ngh]--
			if inDegree[ngh] == 0 {
				queue = append(queue, ngh)
				res = append(res, ngh)
			}
		}
	}

	if len(res) != numCourses {
		return []int{}
	}
	return res
}
