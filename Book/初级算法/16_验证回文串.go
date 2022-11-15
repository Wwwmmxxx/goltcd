package main

import (
	"strings"
)

/**

如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
字母和数字都属于字母数字字符。
给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

示例 1:
	输入: s = "A man, a plan, a canal: Panama"
	输出: true
	解释: "amanaplanacanalpanama" 是回文串。

示例 2:
	输入: s = "race a car"
	输出: false
	解释: "raceacar" 不是回文串。

示例 3：
	输入：s = " "
	输出：true
	解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。

由于空字符串正着反着读都一样，所以是回文串。

提示：
	1 <= s.length <= 2 * 105
	s 仅由可打印的 ASCII 字符组成

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xne8id/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

func isPalindrome(s string) bool {

	s1 := strings.ToUpper(s)
	left, right := 0, len(s1)-1

	for left < right {

		// if !((s1[left] <= 'Z' && s1[left] >= 'A') || (s1[left] >= '0' || s1[left] <= '9')) {
		if !isValid(s[left]) {
			left++
			continue
		}

		// if !((s1[right] <= 'Z' && s1[right] >= 'A') || (s1[right] >= '0' || s1[right] <= '9')) {
		if !isValid(s[right]) {
			right--
			continue
		}

		if s1[right] != s1[left] {
			return false
		}

		right--
		left++
	}

	return true
}

func isValid(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
