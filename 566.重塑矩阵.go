/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] 重塑矩阵
 */

// @lc code=start
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m*n != r*c {
		return mat
	}
	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		result[i/c][i%c] = mat[i/n][i%n]
	}
	return result
}

// @lc code=end
