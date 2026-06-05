// https://leetcode.com/problems/total-waviness-of-numbers-in-range-i/description/
package totalwavinessofnumbersinrangei

func totalWaviness(num1 int, num2 int) int {
	var res int

	var inner func(num int) int
	inner = func(num int) int {
		if num < 100 {
			return 0
		}

		next := num % 10
		mid := num % 100 / 10
		prev := num % 1000 / 100

		if (mid > next && mid > prev) || (mid < next && mid < prev) {
			return 1 + inner(num/10)
		}

		return inner(num / 10)
	}

	for i := num1; i <= num2; i++ {
		res += inner(i)
	}

	return res
}
