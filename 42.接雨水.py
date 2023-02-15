#
# @lc app=leetcode.cn id=42 lang=python3
#
# [42] 接雨水
#

# @lc code=start
class Solution:
    def trap(self, height: List[int]) -> int:
        ans = 0
        stack = [] #整个栈是单调栈 单调递减的
        for i, h in enumerate(height):
            while stack and height[stack[-1]] < h:
                top = stack.pop()
                # 前面起码有两个元素才能闭合
                if not stack:
                    break
                left = stack[-1]
                currWidth = i - left - 1
                currHeight = min(height[left], height[i]) - height[top]
                ans += currWidth * currHeight
            stack.append(i)
        return ans


# @lc code=end

