/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}

// @lc code=end

