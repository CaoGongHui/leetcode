/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode) {
	var dfs func(*TreeNode)
	pre := &TreeNode{}
	first := &TreeNode{}
	second := &TreeNode{}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != nil && pre.Val > node.Val {
			if first.Val == 0 {
				first = pre
			}
			second = node
		}
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	first.Val, second.Val = second.Val, first.Val
}

// @lc code=end
