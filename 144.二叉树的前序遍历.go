/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	result = append(result, root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	return result
}
// @lc code=end
