// https://leetcode.com/problems/power-of-two/description/
package poweroftwo

// With loop
// func isPowerOfTwo(n int) bool {
// 	if n == 1 {
// 		return true
// 	}
// 	for i := 2; i <= n; i = i << 1 {
// 		if i == n {
// 			return true
// 		}
// 	}
// 	return false
// }

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
