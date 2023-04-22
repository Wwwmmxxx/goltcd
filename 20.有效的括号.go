package goltcd

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	每个右括号都有一个对应的相同类型的左括号。

示例 1：
	输入：s = "()"
	输出：true

示例 2：
	输入：s = "()[]{}"
	输出：true

示例 3：
	输入：s = "(]"
	输出：false
*/

// @lc code=start
func isValid(s string) bool {

	var outArray []byte

	if len(s) == 0 {
		return false
	}

	for _, b := range []byte(s) {
		switch b {
		case '(':
			outArray = append(outArray, ')')
		case '[':
			outArray = append(outArray, ']')
		case '{':
			outArray = append(outArray, '}')
		default:
			if len(outArray) > 0 && outArray[len(outArray)-1] == b {
				outArray = outArray[0 : len(outArray)-1]
				continue
			}
			return false
		}

	}

	return len(outArray) == 0

}

// @lc code=end
