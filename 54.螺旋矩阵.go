/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 0 {
		return []int{}
	}
	row := len(matrix)
	col := len(matrix[0])
	result := make([]int, row*col)
	index := 0
	left, top, right, bottom := 0, 0, col-1, row-1

	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			result[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			result[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for col := right - 1; col >= left; col-- {
				result[index] = matrix[bottom][col]
				index++
			}
			for row := bottom - 1; row > top; row-- {
				result[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}

// @lc code=end
