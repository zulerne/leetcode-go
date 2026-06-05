// https://leetcode.com/problems/remove-element/description/
package removeelement

func removeElement(nums []int, val int) int {
	i, j := 0, 0

	for ; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
