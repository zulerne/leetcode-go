// https://leetcode.com/problems/course-schedule-ii/description/
package coursescheduleii

func findOrder(numCourses int, prerequisites [][]int) []int {
	var res []int
	var finished int
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		inDegree[p[0]]++
	}

	var queue []int
	for c := range numCourses {
		if inDegree[c] == 0 {
			queue = append(queue, c)
			res = append(res, c)
			finished++
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
				finished++
			}
		}
	}

	if finished != numCourses {
		return []int{}
	}
	return res
}
