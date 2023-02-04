#
# @lc app=leetcode.cn id=53 lang=python3
#
# [53] 最大子数组和
#

# @lc code=start
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        array = [0] * len(nums)
        array[0] = nums[0]
        for i in range(1, len(nums)):
            array[i] = max(nums[i], nums[i]+array[i-1])
        return max(array)


# @lc code=end
