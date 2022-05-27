/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummy := &ListNode{Next: head}
	for sublength := 1; sublength < length; sublength <<= 1 {
		prev, cur := dummy, dummy.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < sublength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < sublength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			prev.Next = mergeList(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummy.Next
}
func mergeList(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	node := &ListNode{Val: -1}
	root := node
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			node.Next = head1
			head1 = head1.Next
		} else {
			node.Next = head2
			head2 = head2.Next
		}
		node = node.Next
	}
	if head1 != nil {
		node.Next = head1
	}
	if head2 != nil {
		node.Next = head2
	}
	return root.Next
}

// @lc code=end
