// Package kidswithcandies
// https: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies
package kidswithcandies

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	maxCandies := slices.Max(candies)

	for i, candy := range candies {
		res[i] = candy+extraCandies >= maxCandies
	}

	return res
}
