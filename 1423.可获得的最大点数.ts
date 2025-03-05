/*
 * @lc app=leetcode.cn id=1423 lang=typescript
 *
 * [1423] 可获得的最大点数
 */




// @lc code=start
function maxScore(cardPoints: number[], k: number): number {

    // 逆向思维
    // let sum = 0;
    // let total = 0;
    // for (let i = 0; i < cardPoints.length - k; i++) {
    //     sum += cardPoints[i];
    //     total += cardPoints[i]
    // }
    // let res = sum;
    // for (let i = cardPoints.length - k; i < cardPoints.length; i++) {
    //     const element = cardPoints[i];
    //     total += element;
    //     sum += element - cardPoints[i - cardPoints.length + k];
    //     res = Math.min(res, sum);
    // }
    // return total - res;
    //正向
    // 前 k 个数的和。
    // 前 k−1 个数以及后 1 个数的和。
    // 前 k−2 个数以及后 2 个数的和。
    // ……
    // 前 2 个数以及后 k−2 个数的和。
    // 前 1 个数以及后 k−1 个数的和。
    // 后 k 个数的和。
    let res = 0;
    let sum = 0;
    for (let index = 0; index < k; index++) {
        const element = cardPoints[index];
        sum += element;
    }
    res = sum;
    for (let index = 0; index < k; index++) {
        sum += cardPoints[cardPoints.length - 1 - index] - cardPoints[k - index - 1]
        res = Math.max(res, sum)
    }
    return res;
};
// @lc code=end

