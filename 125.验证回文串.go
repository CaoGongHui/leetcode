/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	b := []byte(s)
	for i := 0, j:=len(b)-1; i <=j;  {
		if unicode.IsPunct(b[i]) {
			i++
			continue
		}
		if unicode.IsPunct(b[j]) {
			j--
			continue
		}
		if b[i] != b[j] {
			return false
		}
		i++
		j--
	}
	return true
}
// @lc code=end
