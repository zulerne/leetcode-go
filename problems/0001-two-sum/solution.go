// https://leetcode.com/problems/two-sum/description/
package twosum

func twoSum(nums []int, target int) []int {
	numIndex := make(map[int]int, len(nums))

	for i, num := range nums {
		diff := target - num

		if j, ok := numIndex[diff]; ok {
			return []int{i, j}
		}

		numIndex[num] = i
	}

	return nil
}
