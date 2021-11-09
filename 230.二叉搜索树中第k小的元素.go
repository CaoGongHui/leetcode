/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	leftCount := count(root.Left)
	if leftCount+1 == k {
		return root.Val
	} else if leftCount >= k {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-leftCount-1)
	}
}
func count(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return count(node.Left) + count(node.Right) + 1
}

// @lc code=end

// func kthSmallest(root *TreeNode, k int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	res := 0
// 	var dfs func(*TreeNode)
// 	dfs = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		dfs(node.Left)
// 		k--
// 		if k == 0 {
// 			res = node.Val
// 			return
// 		}
// 		dfs(node.Right)
// 	}
// 	dfs(root)
// 	return res
// }
