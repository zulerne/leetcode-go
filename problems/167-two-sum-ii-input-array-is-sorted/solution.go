// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum > target {
			r--
		} else {
			l++
		}
	}

	return numbers
}
