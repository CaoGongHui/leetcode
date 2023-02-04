#
# @lc app=leetcode.cn id=349 lang=python3
#
# [349] 两个数组的交集
#

# @lc code=start
from typing import List


class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        # hash表方法
        # hash = dict()
        # res = []
        # for _, value in enumerate(nums1):
        #     hash[value] = True

        # for _, value in enumerate(nums2):
        #     if value in hash and hash[value] == True:
        #         hash[value] = False
        #         res.append(value)
        # return res
        s1 = set(nums1)
        s2 = set(nums2)
        return self.intersection_set(s1, s2)
    def intersection_set(self, set1, set2):
        if len(set1) > len(set2):
            set1, set2 = set2, set1
        return [x for x in set1 if x in set2]
# @lc code=end

