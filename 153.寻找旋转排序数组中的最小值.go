/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		mid := left + (right-left)/2
		// if nums[mid] > nums[left] {
		// 	if nums[left] > nums[right] {
		// 		left = mid + 1
		// 	}else {
		// 		right = mid
		// 	}
		// }else {
		// 	if nums[mid] < nums[mid+1] {
		// 		right = mid
		// 	}else{
		// 		left = mid + 1
		// 	}
		// }
		// 跟最后一个元素比 更合理易懂
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

// @lc code=end
