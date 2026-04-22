// https://leetcode.com/problems/max-consecutive-ones-iii/description/
package maxconsecutiveonesiii

func longestOnes(nums []int, k int) int {
	res := 0
	zeros := 0
	begin := 0
	for end := range nums {
		if nums[end] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[begin] == 0 {
				zeros--
			}
			begin++
		}
		res = max(res, end-begin+1)
	}

	return res
}
