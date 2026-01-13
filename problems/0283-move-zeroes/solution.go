// Package movezeroes
// https://leetcode.com/problems/move-zeroes/?envType=study-plan-v2&envId=leetcode-75
package movezeroes

func moveZeroes(nums []int) {
	write := 0

	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write] = nums[read]
			write++
		}
	}

	for i := write; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroesSwap(nums []int) {
	write := 0
	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write], nums[read] = nums[read], nums[write]
			write++
		}
	}
}
