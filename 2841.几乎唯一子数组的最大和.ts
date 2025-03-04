/*
 * @lc app=leetcode.cn id=2841 lang=typescript
 *
 * [2841] 几乎唯一子数组的最大和
 */

// @lc code=start
function maxSum(nums: number[], m: number, k: number): number {
    let map = new Map<number, number>()
    let result = 0
    let sum = 0
    for (let index = 0; index < nums.length; index++) {
        const element = nums[index];
        sum += element
        map.set(element, (map.get(element)?? 0) + 1)
        if (index < k - 1) {
            continue
        }
        if (map.size >= m) {
            result = Math.max(result, sum)
        }
        let out = nums[index - k + 1]
        sum -= out
        if (map.get(out) === 1) {
            map.delete(out)
        } else if (map.has(out)) {
            map.set(out, (map.get(out)?? - 1))
        }

    }
    return result;
};
// @lc code=end

