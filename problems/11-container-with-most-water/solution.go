// Package watercontainer
// https://leetcode.com/problems/container-with-most-water
package watercontainer

// O(N^2)
func maxAreaGreedy(height []int) int {
	area := 0

	for i := 0; i < len(height)-1; i++ {
		h1 := height[i]

		for j := i + 1; j < len(height); j++ {
			h2 := height[j]

			width := j - i

			a := width * min(h1, h2)
			if a > area {
				area = a
			}
		}
	}

	return area
}

// O(N)
func maxArea(height []int) int {
	area := 0

	for left, right := 0, len(height)-1; left < right; {
		hLeft, hRight := height[left], height[right]
		curArea := (right - left) * min(hLeft, hRight)

		if curArea > area {
			area = curArea
		}

		if hLeft > hRight {
			right--
		} else {
			left++
		}
	}

	return area
}
