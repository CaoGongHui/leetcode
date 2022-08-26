/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	stack := []byte{}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	for i := 0; i < length; i++ {
		if pair[s[i]] == 0 {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pair[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

// @lc code=end

