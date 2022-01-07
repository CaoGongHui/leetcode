/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	i := 0
	for k, v := range nums {
		if v != 0 {
			nums[i], nums[k] = v, nums[i]
			i++
		}
	}
	// for j:=i; j<len(nums); j++ {
	// 	nums[j] = 0
	// }
}

// @lc code=end
