/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	// 通过栈解决
	// if len(tokens) == 0 {
	// 	return 0
	// }
	// stack := make([]int, 0)
	// for _, v := range tokens {
	// 	val, err := strconv.Atoi(v)
	// 	if err == nil {
	// 		stack = append(stack, val)
	// 	} else {
	// 		num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
	// 		stack = stack[:len(stack)-2]
	// 		switch v {
	// 		case "+":
	// 			stack = append(stack, num1+num2)
	// 		case "-":
	// 			stack = append(stack, num1-num2)
	// 		case "*":
	// 			stack = append(stack, num1*num2)
	// 		default:
	// 			stack = append(stack, num1/num2)
	// 		}
	// 	}
	// }
	// return stack[0]
	//通过数组模拟栈
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, (len(tokens)+1)/2)
	index := -1
	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err == nil {
			index++
			stack[index] = val
		} else {
			index--
			switch v {
			case "+":
				stack[index] += stack[index+1]
			case "-":
				stack[index] -= stack[index+1]
			case "*":
				stack[index] *= stack[index+1]
			default:
				stack[index] /= stack[index+1]
			}
		}
	}
	return stack[0]
}

// @lc code=end
