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
        slow, fast = l1, l2
        dummy = ListNode(0)
        pre = dummy

        while slow and fast:
            if slow.val < fast.val:
                dummy.next = slow
                slow = slow.next
            else:
                dummy.next = fast
                fast = fast.next
            dummy = dummy.next
        if slow:
            dummy.next = slow
        if fast:
            dummy.next = fast
        return pre.next


# @lc code=end
