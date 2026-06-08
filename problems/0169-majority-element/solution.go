// https://leetcode.com/problems/majority-element/description/
package majorityelement

func majorityElement(nums []int) int {
	var count, cand int

	for _, num := range nums {
		if count == 0 {
			cand = num
			count = 1
		} else if cand == num {
			count++
		} else {
			count--
		}
	}

	return cand
}
