/*
 * @lc app=leetcode.cn id=1652 lang=golang
 *
 * [1652] 拆炸弹
 */

// @lc code=start
func decrypt(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}
	// ans := make([]int, len(code))
	// sum := 0
	// length := len(code)
	// r := k + 1
	// if k < 0 {
	// 	r = length
	// 	k = -k
	// }
	// for _, v := range code[r-k : r] {
	// 	sum += v
	// }
	// for i := range ans {
	// 	ans[i] = sum
	// 	sum += code[r%length] - code[(r-k)%length]
	// 	r++
	// }
	// return ans
	n := len(code)
	ans := make([]int, n)
	code = append(code, code...)
	l, r := 1, k
	if k < 0 {
		l = n + k
		r = n - 1
	}
	sum := 0
	for _, v := range code[l : r+1] {
		sum += v
	}
	for i := range ans {
		ans[i] = sum
		sum += code[r+1] - code[l]
		l++
		r++
	}
	return ans
}

// @lc code=end

