#
# @lc app=leetcode.cn id=1 lang=python3
#
# [1] 两数之和
#

# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashtable = dict()
        for key, value in enumerate(nums):
            if target - value in hashtable:
                return [hashtable[target - value], key]
            hashtable[value] = key
        return []
# @lc code=end

