#
# @lc app=leetcode.cn id=121 lang=python3
#
# [121] 买卖股票的最佳时机
#

# @lc code=start
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max = 0
        min = prices[0]
        for i, val in enumerate(prices):
            if prices[i] < min:
                min = prices[i]
            else:
                dif = prices[i] - min
                if dif > max:
                    max = dif
        return max
# @lc code=end

