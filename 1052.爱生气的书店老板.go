/*
 * @lc app=leetcode.cn id=1052 lang=golang
 *
 * [1052] 爱生气的书店老板
 */

// @lc code=start
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    s := [2]int{}
	res := 0
	for k, v := range customers {
		s[grumpy[k]] += v
		if k < minutes - 1 {
			continue
		}
		res = max(res, s[1])
		if grumpy[k-minutes +1] == 1 {
			s[1] -= customers[k-minutes+1]
		}
	}
	return res + s[0]
}
// @lc code=end

