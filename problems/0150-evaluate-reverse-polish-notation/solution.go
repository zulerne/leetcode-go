// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
package evaluatereversepolishnotation

import "strconv"

func evalRPN(tokens []string) int {
	ops := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "-", "+", "*", "/":
			a, b := ops[len(ops)-2], ops[len(ops)-1]
			ops = ops[:len(ops)-2]

			switch token {
			case "-":
				ops = append(ops, a-b)
			case "+":
				ops = append(ops, a+b)
			case "*":
				ops = append(ops, a*b)
			case "/":
				ops = append(ops, a/b)
			}
		default:
			num, _ := strconv.Atoi(token)
			ops = append(ops, num)
		}
	}

	return ops[0]
}
