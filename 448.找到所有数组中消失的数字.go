/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	var result []int
	for i, v := range nums {
		if v <= n {
			result = append(result, i+1)
		}
	}
	return result
}

// @lc code=end
