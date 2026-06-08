// https://leetcode.com/problems/best-time-to-buy-and-sell-stock//description/
package besttimetobuyandsellstock

import "math"

func maxProfit(prices []int) int {
	var res int
	minPrice := math.MaxInt

	for _, p := range prices {
		minPrice = min(minPrice, p)
		res = max(res, p-minPrice)
	}

	return res
}
