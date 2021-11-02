/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
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
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	var stack []*TreeNode
	var m = make(map[int]int)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := m[k-root.Val]; ok {
			return true
		}
		m[root.Val] = 1
		root = root.Right
	}
	return false
}
// @lc code=end


// if root == nil {
// 	return false
// }
// var stack []*TreeNode
// var m = make(map[int]bool)
// var dfs func (root *TreeNode)
// dfs = func(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	dfs(root.Left)
// 	stack = append(stack, root)
// 	dfs(root.Right)
// }
// dfs(root)
// for i := 0; i < len(stack); i++ {
// 	if _, ok := m[k-stack[i].Val]; ok {
// 		return true
// 	} else {
// 		m[stack[i].Val] = true
// 	}
// }
// return false