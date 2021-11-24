/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
 */

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode{root}
	for len(queue) > 0 {
		var current []int
		q := queue
		queue = nil
		for _, node := range q {
			current = append(current, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, current)
	}
	return reverse(result)
}
func reverse(result [][]int) [][]int {
	var reversed [][]int
	for i:=len(result)-1; i>=0; i-- {
		reversed = append(reversed, result[i])
	}
	return reversed
}
// @lc code=end
