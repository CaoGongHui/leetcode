/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	// 排序法
	// if len(nums) == 1 {
	// 	return nums[0]
	// }
	// sort.Ints(nums)
	// return nums[len(nums)/2]

	// 哈希表
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

// @lc code=end
