/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
	// k = k % len(nums)
	// reverse(nums, 0, len(nums)-1)
	// reverse(nums, 0, k-1)
	// reverse(nums, k, len(nums)-1)
}

// func reverse(nums []int, start, end int) {
// 	for start < end {
// 		nums[start], nums[end] = nums[end], nums[start]
// 		start++
// 		end--
// 	}
// }

// @lc code=end
