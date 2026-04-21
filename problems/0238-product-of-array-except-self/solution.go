// Package productexceptself
// https://leetcode.com/problems/product-of-array-except-self
package productexceptself

func productExceptSelfWithDivision(nums []int) []int {
	res := make([]int, len(nums))

	totalProduct := 1
	totalZeros := 0
	lastZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if totalZeros >= 2 {
			break
		}
		n := nums[i]
		if n != 0 {
			totalProduct *= n
		} else {
			lastZeroIndex = i
			totalZeros++
		}
	}
	if totalZeros >= 2 {
		return res
	}
	if totalZeros == 1 {
		res[lastZeroIndex] = totalProduct
		return res
	}

	for i := 0; i < len(nums); i++ {
		res[i] = totalProduct / nums[i]
	}

	return res
}

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)

	leftProduct := 1
	for i := 0; i < length; i++ {
		res[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := length - 1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return res
}
