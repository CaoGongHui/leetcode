/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	max := 0
	min := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if dif := prices[i] - min; dif > max {
			max = prices[i] - min
		}
	}
	return max
}

// @lc code=end
