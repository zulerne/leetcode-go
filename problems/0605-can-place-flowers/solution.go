// Package placeflowers
// https://leetcode.com/problems/can-place-flowers/description
package placeflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		emtpyLeft := (i == 0) || (flowerbed[i-1] == 0)
		emptyRight := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)

		if emtpyLeft && emptyRight {
			flowerbed[i] = 1
			n--

			if n == 0 {
				return true
			}
		}
	}

	return false
}
