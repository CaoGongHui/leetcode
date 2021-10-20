/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	var temp []int
	p, q := 0, 0
	for {
		if p == m {
			temp = append(temp, nums2[q:]...)
			break
		}
		if q == n {
			temp = append(temp, nums1[p:]...)
			break
		}
		if nums1[p] > nums2[q] {
			temp = append(temp, nums2[q])
			q++
		} else {
			temp = append(temp, nums1[p])
			p++
		}
	}
	mod := (m + n) % 2
	mid := (m + n) / 2
	var result_value float64
	if mod == 0 {
		result_value = (float64(temp[mid]) + float64(temp[mid-1])) / 2
	} else {
		result_value = float64(temp[mid])
	}
	return result_value
}

// @lc code=end
