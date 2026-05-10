// https://leetcode.com/problems/keys-and-rooms/description/
package keysandrooms

func canVisitAllRooms(rooms [][]int) bool {
	stack := [][]int{rooms[0]}
	count := 1
	seen := make([]bool, len(rooms))

	seen[0] = true

	for len(stack) > 0 {
		keys := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, k := range keys {
			if !seen[k] {
				seen[k] = true
				count++
				stack = append(stack, rooms[k])
			}
		}
	}

	return count == len(rooms)
}
