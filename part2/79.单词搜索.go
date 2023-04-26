package goltcd

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode.cn/problems/word-search/description/
 *
 * algorithms
 * Medium (46.32%)
 * Likes:    1572
 * Dislikes: 0
 * Total Accepted:    418.5K
 * Total Submissions: 903.6K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false
 * 。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCCED"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "SEE"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCB"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n = board[i].length
 * 1
 * 1
 * board 和 word 仅由大小写英文字母组成
 *
 *
 *
 *
 * 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
 *
 */

// @lc code=start
func exist(board [][]byte, word string) bool {

	left, right, top, down := 0, len(board)-1, 0, len(board[0])-1

	var f func(int, int, int) bool
	f = func(lineIndex, columnIndex, wordIndex int) bool {

		if wordIndex == len(word) {
			return true
		}

		if lineIndex < left || lineIndex > right || columnIndex < top || columnIndex > down {
			return false
		}

		// 不能走回头路
		if board[lineIndex][columnIndex] == '*' {
			return false
		}

		if board[lineIndex][columnIndex] != word[wordIndex] {
			return false
		}
		tmp := board[lineIndex][columnIndex]
		board[lineIndex][columnIndex] = '*'
		defer func() {
			board[lineIndex][columnIndex] = tmp
		}()

		result := f(lineIndex-1, columnIndex, wordIndex+1) ||
			f(lineIndex+1, columnIndex, wordIndex+1) ||
			f(lineIndex, columnIndex+1, wordIndex+1) ||
			f(lineIndex, columnIndex-1, wordIndex+1)

		return result
	}

	for lineNo, line := range board {
		for columnNo := range line {
			if f(lineNo, columnNo, 0) {
				return true
			}
		}
	}

	return false
}

// @lc code=end
