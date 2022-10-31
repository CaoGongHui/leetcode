#
# @lc app=leetcode.cn id=14 lang=python3
#
# [14] 最长公共前缀
#

# @lc code=start
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        prefix = strs[0]
        if len(strs) > 0:
            for _, str in enumerate(strs):
                while str.find(prefix) != 0:
                    if len(prefix) == 0:
                        return "" 
                    prefix = prefix[:len(prefix)-1]
        return prefix

# @lc code=end

