/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
// func removeDuplicates(nums []int) int {
// 	if len(nums) <= 2 {
// 		return len(nums)
// 	}
// 	slow, fast := 2, 2
// 	for fast < len(nums) {
// 		if nums[slow-2] != nums[fast] {
// 			nums[slow] = nums[fast]
// 			slow++
// 		}
// 		fast++
// 	}
// 	return slow
// }

func removeDuplicates(nums []int) int {
	i := 0
	for _, v := range nums {
		if i < 2 || v > nums[i-2] {
			nums[i] = v
			i++
		}
	}
	return i
}

// @lc code=end
