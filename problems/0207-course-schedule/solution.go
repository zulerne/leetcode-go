// https://leetcode.com/problems/course-schedule/description/
package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	var finished int
	inDegree := make([]int, numCourses)
	graph := make([][]int, numCourses)

	for _, p := range prerequisites {
		inDegree[p[0]]++
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	queue := []int{}
	for c := range numCourses {
		if inDegree[c] == 0 {
			queue = append(queue, c)
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
				finished++
			}
		}
	}

	return numCourses == finished
}
