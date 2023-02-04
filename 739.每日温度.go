/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	//暴力解法
	// length := len(temperatures)
	// result := make([]int, length)
	// for i:= 0;i<length;i++{
	// 	count := 1
	// 	for j := i+1;j<length;j++{
	// 		if temperatures[j] > temperatures[i] {
	// 			result[i] = count
	// 			break
	// 		}else {
	// 			count += 1
	// 		}
	// 	}
	// }
	//单调栈
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}

// @lc code=end

