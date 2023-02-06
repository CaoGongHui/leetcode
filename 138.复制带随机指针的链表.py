#
# @lc app=leetcode.cn id=138 lang=python3
#
# [138] 复制带随机指针的链表
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if not head:
            return None
        d = dict()
        p = head
        while p:
            tmp = Node(p.val, None, None)
            d[p] = tmp
            p = p.next

        p = head
        while p:
            if p.next:
                d[p].next = d[p.next]
            if p.random:
                d[p].random = d[p.random]
            p = p.next
        return d[head]



        
# @lc code=end

