package book01

/**

给你一个整数数组 prices , 其中 prices[i] 表示某支股票第 i 天的价格.
在每一天, 你可以决定是否购买和/或出售股票. 你在任何时候"最多"只能持有"一股"股票. 你也可以先购买, 然后在"同一天"出售。
返回你能获得的"最大"利润。

示例 1：
	输入: prices = [7,1,5,3,6,4]
	输出: 7
	解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
		随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
		总利润为 4 + 3 = 7 。

示例 2：
	输入：prices = [1,2,3,4,5]
	输出：4
	解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4.
		总利润为 4 。

示例3：
	输入：prices = [7,6,4,3,1]
	输出：0
	解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0 。

提示：
	1 <= prices.length <= 3 * 104
	0 <= prices[i] <= 104

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2zsx1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

func maxProfit1(prices []int) int {

	profit := 0
	left := 0

	for left < len(prices)-1 {

		// 购买
		if prices[left] < prices[left+1] {

			right := left + 1

			// 找什么时候卖出
			for ; right < len(prices)-1; right++ {

				// 说明找到了拐点, 那么j就是我们要卖出的日子
				// 找到拐点有两种情况: 1. 第二天金额 < 当天金额; 2. 最后一天, 没有更多金额了

				if (right + 1) == len(prices) {
					right = right + 1
					break
				}

				if prices[right] > prices[right+1] {
					break
				}
			}

			// 存在拐点的话
			profit += prices[right] - prices[left]
			left = right + 1
		} else {
			left++
		}

	}
	return profit
}

func maxProfit(prices []int) int {

	profit := 0

	for i := 1; i < len(prices); i++ {
		d := prices[i] - prices[i-1]
		if d > 0 {
			profit += d
		}
	}

	return profit
}
