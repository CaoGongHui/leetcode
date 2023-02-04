/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	if len(s) == 0{
		return s
	}
	numStack := []int{}
	strStack := []int{}
	num := 0
	result := ""
	for _, v := range s {
		if v > '0' && v < '9' {
			val, _ := strconv.Atoi(v)
			num = num * 10 + val
		}else if v == '[' {
			strStack = append(strStack, result)
			result = ""
			numStack = append(numStack, num)
			num = 0
		}else if v == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			result = string(str) + strings.Repeat(result, count)
		}else {
			result = result + string(v)
		}
	}
	return result
}
// @lc code=end
