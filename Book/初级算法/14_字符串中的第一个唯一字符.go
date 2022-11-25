package book01

/**

给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。

示例 1:
	输入: s = "leetcode"
	输出: 0

示例 2:
	输入: s = "loveleetcode"
	输出: 2

示例 3:
	输入: s = "aabb"
	输出: -1

提示:
	1. 1 <= s.length <= 105
	2. s 只包含小写字母
*/

func firstUniqChar(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {

	var (
		m = make(map[rune][]int, len(s))
	)

	for i, v := range s {

		value, isExist := m[v]

		if !isExist {
			m[v] = []int{i}
		} else {
			m[v] = append(value, i)
		}

	}

	for _, v := range s {
		if len(m[v]) == 1 {
			return m[v][0]
		}
	}

	return -1
}
