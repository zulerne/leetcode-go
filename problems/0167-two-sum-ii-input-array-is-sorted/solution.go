// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}

	return []int{}
}
