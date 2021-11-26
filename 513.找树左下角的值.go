/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
//  层序遍历
// func findBottomLeftValue(root *TreeNode) int {
// 	if root.Left == nil && root.Right == nil {
// 		return root.Val
// 	}
// 	var queue []*TreeNode
// 	queue = append(queue, root)
// 	var result int
// 	for len(queue) > 0 {
// 		size := len(queue)
// 		result = queue[0].Val
// 		for i := 0; i < size; i++ {
// 			node := queue[0]
// 			queue = queue[1:]
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 		}
// 	}
// 	return result
// }
// 递归
func findBottomLeftValue(root *TreeNode) int {
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		// 先右后左 最后剩下的就是zui左下角的值
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
	}
	return root.Val
}

// @lc code=end
