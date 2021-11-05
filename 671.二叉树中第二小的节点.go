/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
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
func findSecondMinimumValue(root *TreeNode) int {
	result := -1
	rootVal := root.Val
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || node.Val > result && result != -1 {
			return
		}
		if node.Val > rootVal {
			result = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

// @lc code=end
