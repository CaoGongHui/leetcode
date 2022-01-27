/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool)
	for _, v := range nums {
		if _, ok := hash[v]; ok {
			return true
		}
		hash[v] = true
	}
	return false
}

// @lc code=end
