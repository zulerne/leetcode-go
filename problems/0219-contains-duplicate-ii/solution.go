// https://leetcode.com/problems/contains-duplicate-ii/description/
package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	indexes := make(map[int]int, len(nums))

	for i, v := range nums {
		if j, ok := indexes[v]; ok && i-j <= k {
			return true
		}
		indexes[v] = i
	}

	return false
}
