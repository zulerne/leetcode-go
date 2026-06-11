// https://leetcode.com/problems/integer-to-roman/description/
package integertoroman

// func intToRoman(num int) string {
// 	var res []byte
// 	dec := 0
// 	for num != 0 {
// 		last := num % 10
// 		switch last {
// 		case 9:
// 			switch dec {
// 			case 0:
// 				res = append(res, 'X', 'I')
// 			case 1:
// 				res = append(res, 'C', 'X')
// 			case 2:
// 				res = append(res, 'M', 'C')
// 			default:
// 			}
// 		case 4:
// 			switch dec {
// 			case 0:
// 				res = append(res, 'V', 'I')
// 			case 1:
// 				res = append(res, 'L', 'X')
// 			case 2:
// 				res = append(res, 'D', 'C')
// 			default:
// 			}
// 		default:
// 			five, one := byte('V'), byte('I')
// 			switch dec {
// 			case 1:
// 				five, one = 'L', 'X'
// 			case 2:
// 				five, one = 'D', 'C'
// 			case 3:
// 				one = 'M'
// 			default:
// 			}

// 			for range last % 5 {
// 				res = append(res, one)
// 			}
// 			if last >= 5 {
// 				res = append(res, five)
// 			}
// 		}

// 		num /= 10
// 		dec++
// 	}

// 	slices.Reverse(res)
// 	return string(res)
// }

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res []byte
	for i, v := range vals {
		for num >= v {
			res = append(res, syms[i]...)
			num -= v
		}
	}
	return string(res)
}
