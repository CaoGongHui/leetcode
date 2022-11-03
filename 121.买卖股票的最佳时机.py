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
        for _, val in enumerate(prices):
            if val < min:
                min = val
            else:
                if val - min > max:
                    max = val - min
        return max
# @lc code=end

