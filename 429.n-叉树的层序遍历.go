/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	var result [][]int
	if root == nil {
		return nil
	}
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		var current []int
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			current = append(current, node.Val)
			for _, node := range node.Children {
				queue = append(queue, node)
			}
		}
		result = append(result, current)
	}
	return result
}

// @lc code=end
