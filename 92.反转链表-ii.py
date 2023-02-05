#
# @lc app=leetcode.cn id=92 lang=python3
#
# [92] 反转链表 II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:

        #一般反转链表的方法
        def reverse_list(head: ListNode):
            pre = ListNode(0)
            cur = head
            while cur:
                next = cur.next
                cur.next = pre
                pre = cur
                cur = next
        dummy_node = ListNode(0)
        dummy_node.next = head
        pre = dummy_node

        # 找到left节点的前一个节点
        for _ in range(left-1):
            pre = pre.next
        # 找到right节点的
        right_node = pre
        for _ in range(right-left+1):
            right_node = right_node.next
        #记录right节点的后一个节点，以便后续访问
        current = right_node.next
        left_node = pre.next
        #先断开链表
        pre.next =None
        right_node.next = None
        #反转链表
        reverse_list(left_node)
        #left节点的前一个节点应该指向right节点,即是反转后的头节点
        pre.next = right_node
        #left节点,反转后变成最后一个节点,他的下一个要指向之前的right节点的后一个节点
        left_node.next = current

        return dummy_node.next




# @lc code=end

