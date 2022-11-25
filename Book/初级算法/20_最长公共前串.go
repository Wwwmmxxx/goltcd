package main

/**

编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：
	输入：strs = ["flower","flow","flight"]
	输出："fl"

示例 2：
	输入：strs = ["dog","racecar","car"]
	输出：""
	解释：输入不存在公共前缀。

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnmav1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

**/

func longestCommonPrefix(strs []string) string {

	var (
		ptrs = make([]int, 0, len(strs))
	)

	loopTimes := getShortestString(strs)

	for j := 0; j < len(ptrs); j++ {

	}

	return ""
}

func getShortestString(strs []string) int {

	result := len(strs[0])

	for _, v := range strs {
		if len(v) < result {
			result = len(v)
		}
	}

	return result
}
