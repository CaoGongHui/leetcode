#
# @lc app=leetcode.cn id=15 lang=python3
#
# [15] 三数之和
#

# @lc code=start
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        result = list()
        nums.sort()
        n = len(nums)
        for first in range(n-2):
            if first > 0 and nums[first] == nums[first-1]:
                continue
            third = n - 1
            target = -1 * nums[first]
            for second in range(first+1, n-1):
                if second > first + 1 and nums[second] == nums[second-1]:
                    continue
                while second < third and nums[second] + nums[third] > target:
                    third = third - 1
                if second == third:
                    break
                if nums[second] + nums[third] == target:
                    result.append(
                        [nums[first], nums[second], nums[third]])
        return result
        # @lc code=end
