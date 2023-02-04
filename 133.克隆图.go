/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var cg func(*Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if v, ok := visited[node]; ok {
			return v
		}
		newNode := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
		visited[node] = newNode
		for _, v := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, cg(v))
		}
		return newNode
	}
	return cg(node)
}

// @lc code=end

