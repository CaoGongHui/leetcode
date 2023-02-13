#
# @lc app=leetcode.cn id=946 lang=python3
#
# [946] 验证栈序列
#

# @lc code=start
class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        stack, j = [], 0
        for i in pushed:
            stack.append(i)
            while stack and stack[-1] == popped[j]:
                stack.pop()
                j += 1
        return 0 == len(stack)

# @lc code=end
