package book01

/**

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

示例 1:
	输入: s = "anagram", t = "nagaram"
	输出: true

示例 2:
	输入: s = "rat", t = "car"
	输出: false

提示:
	1. 1 <= s.length, t.length <= 5 * 104
	2. s 和 t 仅包含小写字母

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn96us/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

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

func isAnagram1(s string, t string) bool {

	var (
		s1 = make(map[rune]int, len(s))
		t1 = make(map[rune]int, len(t))
	)

	for _, v := range s {
		s1[v]++
	}

	for _, v := range t {
		t1[v]++
	}

	if len(s1) != len(t1) {
		return false
	}

	for key, value := range s1 {
		if t1[key] != value {
			return false
		}
	}

	return true
}
