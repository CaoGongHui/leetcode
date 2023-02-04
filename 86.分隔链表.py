#
# @lc app=leetcode.cn id=86 lang=python3
#
# [86] 分隔链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        low_head = low_tail = ListNode(0)
        high_head = high_tail = ListNode(0)
        while head:
            if head.val < x:
                low_tail.next = head
                low_tail = low_tail.next
            else:
                high_tail.next = head
                high_tail = high_tail.next
            head = head.next
        low_tail.next = high_head.next
        high_tail.next = None
        return low_head.next
# @lc code=end


# @lc code=end

