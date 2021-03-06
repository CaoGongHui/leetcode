/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

// func lengthOfLongestSubstring(s string) int {
// 	start, end := 0, 0
// 	for i := 0; i < len(s); i++ {
// 		index := strings.Index(s[start:i], string(s[i]))
// 		if index == -1 {
// 			if i+1 > end {
// 				end = i + 1
// 			}
// 		} else {
// 			start += index + 1
// 			end += index + 1
// 		}
// 	}
// 	return end - start
// }

// @lc code=end
