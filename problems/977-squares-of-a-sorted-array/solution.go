// https://leetcode.com/problems/squares-of-a-sorted-array/description/
package squaresofasortedarray

func sortedSquares(nums []int) []int {
	res, resIdx := make([]int, len(nums)), len(nums)-1
	l, r := 0, len(nums)-1

	for ; l <= r; resIdx-- {
		vl, vr := nums[l]*nums[l], nums[r]*nums[r]
		if vl < vr {
			res[resIdx] = vr
			r--
		} else {
			res[resIdx] = vl
			l++
		}
	}

	return res
}
