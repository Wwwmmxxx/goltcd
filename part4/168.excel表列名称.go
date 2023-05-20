package goltcd

/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 *
 * https://leetcode.cn/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (43.84%)
 * Likes:    614
 * Dislikes: 0
 * Total Accepted:    133.4K
 * Total Submissions: 304.2K
 * Testcase Example:  '1'
 *
 * 给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
 *
 * 例如：
 *
 *
 * A -> 1
 * B -> 2
 * C -> 3
 * ...
 * Z -> 26
 * AA -> 27
 * AB -> 28
 * ...
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：columnNumber = 1
 * 输出："A"
 *
 *
 * 示例 2：
 *
 *
 * 输入：columnNumber = 28
 * 输出："AB"
 *
 *
 * 示例 3：
 *
 *
 * 输入：columnNumber = 701
 * 输出："ZY"
 *
 *
 * 示例 4：
 *
 *
 * 输入：columnNumber = 2147483647
 * 输出："FXSHRXW"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

// @lc code=start
// 这道题实际上是表示为26进制, 即每到Z+1就要进位1, 变成AA
// 反过来说, 若到AA-1, 则变成Z
// AA => 26 ^ 1 + 26 ^ 0 * 1
// AB => 26 ^ 1 + 26 ^ 0 * 2
// columnNumber => 结果, 实际是10进制转26进制
func convertToTitle(columnNumber int) string {
	m := map[int]byte{
		1: 'A', 2: 'B', 3: 'C', 4: 'D', 5: 'E',
		6: 'F', 7: 'G', 8: 'H', 9: 'I', 10: 'J',
		11: 'K', 12: 'L', 13: 'M', 14: 'N', 15: 'O',
		16: 'P', 17: 'Q', 18: 'R', 19: 'S', 20: 'T',
		21: 'U', 22: 'V', 23: 'W', 24: 'X', 25: 'Y',
		26: 'Z',
	}
	var res []byte

	for columnNumber > 0 {
		k := columnNumber % 26
		// 碰到0, 要退位, 因为%26==0的情况是26, 0-25对应的是A-Z, 26就要退一位
		// 打个比方2进制 %2==0的情况是2%2==0, 碰到2就要进位, 同理%26
		if k == 0 {
			k = 26
			columnNumber -= 26
		}
		res = append([]byte{m[k]}, res...)
		columnNumber = columnNumber / 26
	}

	return string(res)
}

// @lc code=end
