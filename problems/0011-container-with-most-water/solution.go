// https://leetcode.com/problems/container-with-most-water/description/
package containerwithmostwater

func maxArea(height []int) int {
	var res int
	begin, end := 0, len(height)-1

	for begin < end {
		width := end - begin
		if height[begin] > height[end] {
			res = max(res, height[end]*width)
			end--
		} else {
			res = max(res, height[begin]*width)
			begin++
		}
	}

	return res
}
