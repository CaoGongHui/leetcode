/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	// 找中间节点
	mid := findMiddle(head)
	// 翻转中间节点后面的链表并从中间节点断开链接
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	// 合并两个链表
	mergeList(l1, l2)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func reverseList(node *ListNode) *ListNode {
	var prev, cur *ListNode = nil, node
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
func mergeList(l1, l2 *ListNode) {
	var l1tmp, l2tmp *ListNode
	for l1 != nil && l2 != nil {
		l1tmp = l1.Next
		l2tmp = l2.Next
		l1.Next = l2
		l1 = l1tmp
		l2.Next = l1
		l2 = l2tmp
	}
}

// @lc code=end
