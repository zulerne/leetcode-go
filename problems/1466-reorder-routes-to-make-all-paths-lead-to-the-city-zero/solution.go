// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/description/
package reorderroutestomakeallpathsleadtothecityzero

import "slices"

func minReorder(n int, connections [][]int) int {
	var res int
	graphFull := make([][]int, n)
	graphOrigin := make([][]int, n)
	seen := make([]bool, n)

	for _, conn := range connections {
		graphFull[conn[0]] = append(graphFull[conn[0]], conn[1])
		graphOrigin[conn[0]] = append(graphOrigin[conn[0]], conn[1])
		graphFull[conn[1]] = append(graphFull[conn[1]], conn[0])
	}

	queue := []int{0}
	for len(queue) > 0 {
		city := queue[0]
		queue = queue[1:]

		seen[city] = true
		for _, neigh := range graphFull[city] {
			if !seen[neigh] {
				if !slices.Contains(graphOrigin[neigh], city) {
					res++
				}
				seen[neigh] = true
				queue = append(queue, neigh)
			}
		}
	}

	return res
}

/**
func minReorder(n int, connections [][]int) int {
	var res int
	graphFull := make([][]int, n)
	origin := make(map[[2]int]bool)
	seen := make([]bool, n)

	for _, conn := range connections {
		graphFull[conn[0]] = append(graphFull[conn[0]], conn[1])
		origin[[2]int{conn[0], conn[1]}] = true
		graphFull[conn[1]] = append(graphFull[conn[1]], conn[0])
	}

	queue := []int{0}
	for len(queue) > 0 {
		city := queue[0]
		queue = queue[1:]

		seen[city] = true
		for _, n := range graphFull[city] {
			if !seen[n] {
				if _, ok := origin[[2]int{n, city}]; !ok {
					res++
				}
				seen[n] = true
				queue = append(queue, n)
			}
		}
	}

	return res
}
*/

/**
 func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
 }

 func minReorder(n int, connections [][]int) int {
	var res int
	graphFull := make([][]int, n)
	seen := make([]bool, n)

	for _, conn := range connections {
		graphFull[conn[0]] = append(graphFull[conn[0]], conn[1])
		graphFull[conn[1]] = append(graphFull[conn[1]], -conn[0])
	}

	queue := []int{0}
	for len(queue) > 0 {
		city := queue[0]
		queue = queue[1:]

		seen[city] = true
		for _, neigh := range graphFull[city] {
			absNeigh := abs(neigh)
			if !seen[absNeigh] {
				if neigh > 0 {
					res++
				}
				seen[absNeigh] = true
				queue = append(queue, absNeigh)
			}
		}
	}

	return res
 }
*/
