#
# @lc app=leetcode.cn id=20 lang=python3
#
# [20] 有效的括号
#

# @lc code=start
class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) % 2!= 0:
            return False

        pair = {
            ")": "(",
            "]": "[",
            "}": "{",
        }
        stack = []
        for char in s:
            if char in pair:
                if not stack or stack[-1] != pair[char]:
                    return False
                stack.pop()
            else:
                stack.append(char)
        return len(stack) == 0

# @lc code=end

