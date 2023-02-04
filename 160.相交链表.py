#
# @lc app=leetcode.cn id=160 lang=python3
#
# [160] 相交链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        h1, h2 = headA, headB
        # 双指针法
        # while l1!= l2:
        #     if l1:
        #         l1 = l1.next
        #     else:
        #         l1 = headB
        #     if l2:
        #         l2 = l2.next
        #     else:
        #         l2 = headA
        # return l1
        # hashmap法
        d = {}
        while h1:
            d[h1] = True
            h1 = h1.next
        while h2:
            if h2 in d:
                return h2
            h2 = h2.next
        return None


# @lc code=end
