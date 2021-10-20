/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	result := [][]int{}
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		if len(path) == k {
			temp:= make([]int,k)
			copy(temp,path)
			result = append(result,temp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path,i)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(1, []int{})
	return result
}
// @lc code=end