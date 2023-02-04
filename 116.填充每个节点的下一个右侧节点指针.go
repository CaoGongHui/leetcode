/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}
// 	queue := []*Node{root}
// 	for len(queue) > 0 {
// 		size := len(queue)
// 		for i := 0; i < size; i++ {
// 			node := queue[0]
// 			queue = queue[1:]
// 			if i < size-1 {
// 				node.Next = queue[0]
// 			} else {
// 				node.Next = nil
// 			}
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 		}
// 	}
// 	return root
// }

func connect(root *Node) *Node {
	// if root == nil {
	// 	return nil
	// }
	// queue := []*Node{root}
	// for len(queue) > 0 {
	// 	tmp := queue
	// 	size := len(tmp)
	// 	queue = nil
	// 	for i, node := range tmp {
	// 		if i+1 < size {
	// 			node.Next = tmp[i+1]
	// 		}
	// 		if node.Left != nil {
	// 			queue = append(queue, node.Left)
	// 		}
	// 		if node.Right != nil {
	// 			queue = append(queue, node.Right)
	// 		}
	// 	}
	// }
	// return root
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		tmp := queue
		size := len(tmp)
		queue = nil
		for i, node := range tmp {
			if i+1 < size {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}
// 	for leftMost := root; leftMost.Left != nil; leftMost = leftMost.Left {
// 		for node := leftMost; node != nil; node = node.Next {

// 			node.Left.Next = node.Right
// 			if node.Next != nil {
// 				node.Right.Next = node.Next.Left
// 			}
// 		}
// 	}
// 	return root
// }

// @lc code=end
