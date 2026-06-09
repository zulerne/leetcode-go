// https://leetcode.com/problems/jump-game-ii/description/
package jumpgameii

func jump(nums []int) int {
	var res, curEnd, farthest int
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == curEnd {
			res++
			curEnd = farthest
		}
	}
	return res
}
