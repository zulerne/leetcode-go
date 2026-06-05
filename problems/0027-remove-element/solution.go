// https://leetcode.com/problems/remove-element/description/
package removeelement

func removeElement(nums []int, val int) int {
	first, second := 0, 0

	for ; first < len(nums); first++ {
		if nums[first] != val {
			nums[second] = nums[first]
			second++
		}
	}

	return second
}
