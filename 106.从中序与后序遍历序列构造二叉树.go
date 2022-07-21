/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	indeMap := map[int]int{}
	for i, v := range inorder {
		indeMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}
		indexRoot := indeMap[val]
		root.Right = build(indexRoot+1, right)
		root.Left = build(left, indexRoot-1)
		return root		
	}
	return build(0, len(inorder)-1)
}
// @lc code=end

