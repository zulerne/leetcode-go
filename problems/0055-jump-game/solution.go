// https://leetcode.com/problems/jump-game/description/
package jumpgame

func canJump(nums []int) bool {
	maxReach := 0
	for i, v := range nums {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+v)
	}
	return true
}
