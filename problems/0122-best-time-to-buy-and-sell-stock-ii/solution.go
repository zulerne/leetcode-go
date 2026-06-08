// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
package besttimetobuyandsellstockii

func maxProfit(prices []int) int {
	var res int

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}
