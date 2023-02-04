/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
var (
	dx = [4]int{0, 1, 0, -1}
	dy = [4]int{1, 0, -1, 0}
)

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := dfs(grid, i, j, m, n)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
func dfs(grid [][]int, i, j, m, n int) int {
	count := 1
	grid[i][j] = 0
	for k := 0; k < 4; k++ {
		mx, my := i+dx[k], j+dy[k]
		if mx >= 0 && mx < m && my >= 0 && my < n && grid[mx][my] == 1 {
			count += dfs(grid, mx, my, m, n)
		}
	}
	return count
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

