package goltcd

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 *
 * https://leetcode.cn/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (62.79%)
 * Likes:    856
 * Dislikes: 0
 * Total Accepted:    242.1K
 * Total Submissions: 384.9K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
 * 输出：[[1,0,1],[0,0,0],[1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
 * 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[0].length
 * 1 <= m, n <= 200
 * -2^31 <= matrix[i][j] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
 * 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
 * 你能想出一个仅使用常量空间的解决方案吗？
 *
 *
 */

// @lc code=start
func setZeroes(matrix [][]int) {

	var (
		rowMax, columnMax = len(matrix), len(matrix[0])

		top, down, left, right = 1, 2, 3, 4

		m = make([]int, 0)
	)

	var turnIntoZero func(int, int, int)

	turnIntoZero = func(i, j, description int) {

		if i < 0 || i >= rowMax || j < 0 || j >= columnMax {
			return
		}

		m = append(m, i, j)

		switch description {
		case top:
			turnIntoZero(i-1, j, description)
		case down:
			turnIntoZero(i+1, j, description)
		case left:
			turnIntoZero(i, j-1, description)
		case right:
			turnIntoZero(i, j+1, description)
		}

	}

	for i := 0; i < rowMax; i++ {
		for j := 0; j < columnMax; j++ {
			if matrix[i][j] == 0 {
				turnIntoZero(i-1, j, top)
				turnIntoZero(i+1, j, down)
				turnIntoZero(i, j-1, left)
				turnIntoZero(i, j+1, right)
			}
		}
	}

	for i := 0; i < len(m); i = i + 2 {
		matrix[m[i]][m[i+1]] = 0
	}
}

// @lc code=end
