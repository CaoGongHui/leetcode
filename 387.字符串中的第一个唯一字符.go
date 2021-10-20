/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	arr := make([]int,26)
	s_byte := []byte(s)
	for i:= 0; i < len(s_byte); i++ {
		arr[s_byte[i] - 'a']++
	}
	for i:= 0; i < len(s_byte); i++ {
		if arr[s_byte[i] - 'a'] == 1{
			return i
		}
	}
	return -1
}
// @lc code=end
