/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	s_list := []byte(s)	//不需要转成byte切片 go可以直接遍历字符串
// 	t_list := []byte(t)
// 	hash := make(map[byte]int)
// 	for _, v := range s_list {
// 		if _, ok := hash[v]; !ok {
// 			hash[v] = 1
// 		} else {
// 			hash[v] += 1
// 		}
// 	}
// 	for _, v := range t_list {	// 不必再统计直接再原表基础上减
// 		if _, ok := hash[v]; !ok {
// 			return false
// 		} else {
// 			hash[v] -= 1
// 		}
// 	}
// 	for _, v := range hash {
// 		if v != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

//优化后的hash方法
func isAnagram(s, t string) bool {
    if len(s) != len(t) {
        return false
    }
    cnt := map[rune]int{}
    for _, ch := range s {
        cnt[ch]++
    }
    for _, ch := range t {
        cnt[ch]--
        if cnt[ch] < 0 {
            return false
        }
    }
    return true
}

// @lc code=end
