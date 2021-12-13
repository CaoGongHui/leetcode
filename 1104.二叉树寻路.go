/*
 * @lc app=leetcode.cn id=1104 lang=golang
 *
 * [1104] 二叉树寻路
 */

// @lc code=start
func pathInZigZagTree(label int) []int {
	res := []int{}
	for label > 1 {
		res = append(res, label)
		label = label >> 1
		label = label ^ (1 << (bits.Len(uint(label)) - 1) -1)
	}
	res = append(res, 1)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
// @lc code=end
