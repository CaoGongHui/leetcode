#
# @lc app=leetcode.cn id=206 lang=python3
#
# [206] 反转链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        # 迭代法
        # pre = None
        # cur = head
        # while cur:
        #     next = cur.next
        #     cur.next = pre
        #     pre = cur
        #     cur = next
        # return pre
        # 递归法
        if head is None or head.next is None:
            return head

        p = self.reverseList(head.next)
        head.next.next = head
        head.next = None

        return p


# @lc code=end

# @lc code=end

