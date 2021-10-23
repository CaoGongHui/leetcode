/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	var ans, pre = math.MaxInt64, -1
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if pre != -1 && root.Val-pre < ans {
			ans = root.Val - pre
		}
		pre = root.Val
		inorder(root.Right)
	}
	inorder(root)
	return ans
}

// @lc code=end

