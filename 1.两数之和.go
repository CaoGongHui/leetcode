/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for k, v := range nums {
		if _, ok := hashmap[target-v]; ok {
			return []int{hashmap[target-v], k}
		}
		hashmap[v] = k
	}
	return []int{}
}

// @lc code=end
