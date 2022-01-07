/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	var res []int
	for _, v := range nums {
		if nums[abs(v)-1] < 0 {
			res = append(res, abs(v))
		} else {
			nums[abs(v)-1] = -nums[abs(v)-1]
		}
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
