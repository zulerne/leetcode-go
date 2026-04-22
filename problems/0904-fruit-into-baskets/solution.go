// https://leetcode.com/problems/fruit-into-baskets/description/
package fruitintobaskets

func totalFruit(fruits []int) int {
	res := 0

	count := make([]int, len(fruits))
	types := 0

	begin := 0
	for end, endVal := range fruits {
		if count[endVal] == 0 {
			types++
		}
		count[endVal]++

		for types > 2 {
			beginVal := fruits[begin]
			count[beginVal]--
			if count[beginVal] == 0 {
				types--
			}
			begin++
		}
		res = max(res, end-begin+1)
	}

	return res
}
