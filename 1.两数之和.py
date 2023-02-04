#
# @lc app=leetcode.cn id=1 lang=python3
#
# [1] 两数之和
#

# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashMap = dict()
        for k, v in enumerate(nums):
            if target - v in hashMap:
                return [hashMap[target-v], k]
            hashMap[v] = k
        return []
# @lc code=end
