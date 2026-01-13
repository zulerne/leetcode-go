// Package twosum
// https://leetcode.com/problems/two-sum/description
package twosum

// twoSumFirst has O(n^2) complexity
func twoSumFirst(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// TwoSum finds indices of two numbers that add up to the target.
// Time Complexity: O(n)
// Space Complexity: O(n)
func twoSum(nums []int, target int) []int {
	numIndex := make(map[int]int, len(nums))

	for i, num := range nums {
		diff := target - num

		if j, ok := numIndex[diff]; ok {
			return []int{j, i}
		}

		numIndex[num] = i
	}

	return nil
}
