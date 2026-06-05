// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/
package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if j > 0 && nums[i] == nums[j-1] && nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}

	return j + 1
}
