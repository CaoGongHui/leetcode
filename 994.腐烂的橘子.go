/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start
var (
	di = []int{0, 0, 1, -1}
	dj = []int{1, -1, 0, 0}
)

func orangesRotting(grid [][]int) int {
	var queue [][2]int
	fresh_origin := 0
	col, row := len(grid), len(grid[0])
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if grid[i][j] == 1 {
				fresh_origin++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	if fresh_origin == 0 {
		return 0
	}
	ans := 0
	for len(queue) > 0 {
		ans++
		size := len(queue)
		for i := 0; i < size; i++ {
			tmp := queue[0]
			queue = queue[1:]
			x, y := tmp[0], tmp[1]
			for i := 0; i < 4; i++ {
				next_x, next_y := x+di[i], y+dj[i]
				if next_x >= 0 && next_x < len(grid) && next_y >= 0 && next_y < len(grid[0]) && grid[next_x][next_y] == 1 {
					grid[next_x][next_y] = 2
					fresh_origin--
					queue = append(queue, [2]int{next_x, next_y})
				}
			}
		}
	}
	if fresh_origin > 0 {
		return -1
	}
	return ans - 1
}
