/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{-1,-1}
	}
	left, right := 0, length - 1
	res := []int{-1, -1}
	for left < right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			if nums[mid] > nums[mid-1] {
				res[0] = mid
				left = mid + 1
			}else{
				
			}
		}
		if nums[mid] < target {
			left = mid + 1
		}else {

		}
	}

}
// @lc code=end
