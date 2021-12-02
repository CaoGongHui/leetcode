/*
 * @lc app=leetcode.cn id=623 lang=golang
 *
 * [623] 在二叉树中增加一行
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		left := &TreeNode{Val: val}
		left.Left = root
		return left
	}
	queue := []*TreeNode{root}
	level := 1
	for len(queue) > 0 {
		level++
		q := queue
		queue = nil
		for _, node := range q {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if level == depth {
				ltree := node.Left
				rtree := node.Right
				node.Left = &TreeNode{Val: val}
				node.Left.Left = ltree
				node.Right = &TreeNode{Val: val}
				node.Right.Right = rtree
			}
		}
	}
	return root
}

// @lc code=end
