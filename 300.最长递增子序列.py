#
# @lc app=leetcode.cn id=300 lang=python3
#
# [300] 最长递增子序列
#

# @lc code=start
class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        size = len(nums)
        array = [0] * size
        for i in range(size):
            array[i] = 1
            for j in range(i):
                if nums[i] > nums[j]:
                    array[i] = max(array[j]+1, array[i])
        return max(array)
# @lc code=end
