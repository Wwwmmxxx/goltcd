package goltcd

import "fmt"

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode.cn/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (49.27%)
 * Likes:    1349
 * Dislikes: 0
 * Total Accepted:    357.1K
 * Total Submissions: 724.9K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * -100
 *
 *
 */

// @lc code=start
const (
	right = iota
	down
	left
	up
)

func spiralOrder(matrix [][]int) []int {

	var (
		count  = len(matrix) * len(matrix[0])
		result = make([]int, 0, count)
		m      = make(map[string]bool, count)
		rawMax = len(matrix)
		colMax = len(matrix[0])
	)

	var f func(direction, raw, col, c int)

	f = func(direction, raw, col, c int) {
		result = append(result, matrix[raw][col])
		nextRaw, nextCol := getNextPoint(direction, raw, col)

		// 需要进行转向
		if (direction == right && nextCol == colMax) ||
			(direction == down && nextRaw == rawMax) ||
			(direction == left && nextCol == -1) ||
			m[fmt.Sprintf("%d_%d", nextRaw, nextCol)] {
			direction = (direction + 1) % 4
		}

		m[fmt.Sprintf("%d_%d", raw, col)] = true

		raw, col = getNextPoint(direction, raw, col)
		if c == count {
			return
		}
		f(direction, raw, col, c+1)
	}

	f(right, 0, 0, 1)

	return result
}

func getNextPoint(direction, raw, col int) (int, int) {
	switch direction {
	case right:
		col = col + 1
	case down:
		raw = raw + 1
	case left:
		col = col - 1
	case up:
		raw = raw - 1
	}

	return raw, col
}

func spiralOrder2(matrix [][]int) []int {

	var result []int

	if len(matrix) == 0 {
		return result
	}

	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1

	var x, y int

	for left <= right && up <= down {
		for y = left; y <= right && avoid(left, right, up, down); y++ {
			result = append(result, matrix[x][y])
		}
		y--
		up++

		for x = up; x <= down && avoid(left, right, up, down); x++ {
			result = append(result, matrix[x][y])
		}
		x--
		right--

		for y = right; y >= left && avoid(left, right, up, down); y-- {
			result = append(result, matrix[x][y])
		}
		y++
		down--

		for x = down; x >= up && avoid(left, right, up, down); x-- {
			result = append(result, matrix[x][y])
		}
		x++
		left++
	}

	return result
}

func avoid(left, right, up, down int) bool {
	return up <= down && left <= right
}

// @lc code=end
