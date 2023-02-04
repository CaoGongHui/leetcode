/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := map[int]int{}
	for i, v := range inorder {
		indexMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		val := preorder[0]
		preorder = preorder[1:]
		root := &TreeNode{Val: val}
		indexRoot := indexMap[val]
		root.Left = build(left, indexRoot-1)
		root.Right = build(indexRoot+1, right)
		return root
	}
	return build(0, len(inorder)-1)
}

// @lc code=end

