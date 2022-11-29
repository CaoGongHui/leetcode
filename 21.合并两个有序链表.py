#
# @lc app=leetcode.cn id=21 lang=python3
#
# [21] 合并两个有序链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        slow, fast  = l1, l2
        dummpy = ListNode(0)
        root = dummpy
        while slow != None and fast != None:
            if slow.val > fast.val:
                dummpy.next = fast
                fast = fast.next
            else:
                dummpy.next = slow
                slow = slow.next
            dummpy = dummpy.next
        if slow != None:
            dummpy.next = slow
        if fast != None:
            dummpy.next = fast
        return root.next

# @lc code=end

