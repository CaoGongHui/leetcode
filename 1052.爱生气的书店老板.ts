/*
 * @lc app=leetcode.cn id=1052 lang=typescript
 *
 * [1052] 爱生气的书店老板
 */
// n = 3
// [1,0,1,2,1,1,7,5]
// [0,1,0,1,0,1,0,1]
// @lc code=start
function maxSatisfied(customers: number[], grumpy: number[], minutes: number): number {
    let sum0 = 0
    for (let i = 0; i < grumpy.length; i++) {
        const element = grumpy[i];
        if (element === 0) {
            sum0 += customers[i]
        }
    }
    let res = 0
    let sum1 = 0
    for (let i = 0; i < minutes; i++) {
        const element = grumpy[i];
        if (element === 1) {
            sum1 += customers[i]
        }
    }
    res = sum1
    for (let i = minutes; i < grumpy.length; i++) {
        const element = grumpy[i];
        if (element === 1) {
            sum1 += customers[i]
        }
        if (grumpy[i - minutes] === 1) {
            sum1 -= customers[i - minutes]
        }
        res = Math.max(res, sum1)
    }
    return sum0 + res
};
// @lc code=end

