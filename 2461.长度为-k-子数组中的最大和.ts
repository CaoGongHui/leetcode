/*
 * @lc app=leetcode.cn id=2461 lang=typescript
 *
 * [2461] 长度为 K 子数组中的最大和
 */

// @lc code=start
function maximumSubarraySum(nums: number[], k: number): number {
	let result = 0;
	let map = new Map<number, number>();
	let sum = 0;
	for (let index = 0; index < k; index++) {
		sum += nums[index];
		map.set(nums[index], (map.get(nums[index]) ?? 0) + 1);
	}
	if (map.size === k) {
		result = Math.max(result, sum);
	}
	for (let index = k; index < nums.length; index++) {
		sum += nums[index] - nums[index - k];
		map.set(nums[index], (map.get(nums[index]) ?? 0) + 1);
		map.set(nums[index - k], (map.get(nums[index - k]) ?? 0) - 1);
          if (map.get(nums[index - k]) === 0) {
               map.delete(nums[index - k]);
          }
		if (map.size === k) {
			result = Math.max(result, sum);
		}
	}
	return result;
}
// @lc code=end
