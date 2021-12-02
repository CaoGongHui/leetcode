/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
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
type item struct {
	index int
	*TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result, queue := 1, []item{{1, root}}
	for len(queue) > 0 {
		if width := queue[len(queue)-1].index - queue[0].index + 1; width > result {
			result = width
		}
		var items []item
		for _, obj := range queue {
			if obj.Left != nil {
				items = append(items, item{obj.index * 2, obj.Left})
			}
			if obj.Right != nil {
				items = append(items, item{obj.index*2 + 1, obj.Right})
			}
		}
		queue = items
	}
	return result
}

// @lc code=end
