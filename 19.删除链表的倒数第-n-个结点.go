/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dump := &ListNode{0, head}
	slow := dump
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow := slow.Next
	}
	slow.Next = slow.Next.Next
	return dump.Next
}

// @lc code=end
