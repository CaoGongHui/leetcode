#
# @lc app=leetcode.cn id=66 lang=python3
#
# [66] 加一
#

# @lc code=start
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        add = 0
        i = len(digits) - 1
        while i >= 0:
            digits[i] = digits[i] + 1
            if digits[i] == 10:
                add = 1
                digits[i] = digits[i] % 10
            if add == 1:
                i = i - 1
                add = 0
                if i == -1:
                    digits[0] = digits[0] % 10
                    digits.insert(0, 1)
                    break
            else:
                break
        return digits

# @lc code=end
