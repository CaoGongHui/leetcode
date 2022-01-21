/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start

// func fourSum(nums []int, target int) [][]int {
// 	sort.Ints(nums)
//     n := len(nums)
// 	result := make([][]int, 0)
// 	for first := 0; first < n-3; first++ {
// 		if first > 0 && nums[first] == nums[first-1] {
// 			continue
// 		}
// 		for second := first + 1; second < n-2; second++ {
// 			if second >　first+1 && nums[second] == nums[second-1] {
// 				continue
// 			}
// 			fourth := n-1
// 			tar := target - nums[first] - nums[second]
// 			for third := second + 1; third < n-1; third++ {
// 				if third > second+1 && nums[third] == nums[third-1] {
// 					continue
// 				}
// 				for third < fourth && nums[third]+nums[fourth] > tar {
// 					fourth--
// 				}
// 				if third == fourth {
// 					break
// 				}
// 				if nums[third]+nums[fourth] == tar {
// 					result = append(result, []int{nums[first], nums[second], nums[third], nums[fourth]})
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

func fourSum(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

// @lc code=end
