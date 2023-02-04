#
# @lc app=leetcode.cn id=142 lang=python3
#
# [142] 环形链表 II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:

        slow, fast = head, head

        while fast and fast.next:
            # 要先走 fast走过的路是slow的两倍
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                ptr = head
                while ptr != slow:
                    slow = slow.next
                    ptr = ptr.next
                return ptr
        return None


# @lc code=end

