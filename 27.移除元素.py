#
# @lc app=leetcode.cn id=27 lang=python3
#
# [27] 移除元素
#

# @lc code=start
import enum


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        i = 0
        for _, v in enumerate(nums):
            if v != val:
                nums[i] = v
                i = i + 1
        return i

# @lc code=end
