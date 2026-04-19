// Package template - template for new solutions
// {url}
package palindromnum

func isPalindrome(x int) bool {
	xCopy := x
	reversed := 0
	for xCopy > 0 {
		last := xCopy % 10
		xCopy /= 10
		reversed = reversed*10 + last
	}
	return x == reversed
}
