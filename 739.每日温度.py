#
# @lc app=leetcode.cn id=739 lang=python3
#
# [739] 每日温度
#

# @lc code=start
class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        length = len(temperatures)
        res = [0] * length
        stack = []
        for i in range(length):
            # 栈不为空,且当前温度大于栈顶的元素,栈顶的元素就出栈
            # 并且记录该天的结果
            while stack and temperatures[i] > temperatures[stack[-1]]:
                pos = stack.pop()
                res[pos] = i - pos
            # 将日期入栈
            stack.append(i)
        return res

# @lc code=end
