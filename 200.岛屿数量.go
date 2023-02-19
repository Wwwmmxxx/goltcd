package goLeetCode

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量

给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1:
	输入: grid = [
  		["1","1","1","1","0"],
  		["1","1","0","1","0"],
  		["1","1","0","0","0"],
  		["0","0","0","0","0"]
	]
	输出: 1

示例 2:
	输入: grid = [
  		["1","1","0","0","0"],
  		["1","1","0","0","0"],
  		["0","0","1","0","0"],
  		["0","0","0","1","1"]
	]
	输出：3

提示:
	m == grid.length
	n == grid[i].length
	1 <= m, n <= 300
	grid[i][j] 的值为 '0' 或 '1'
*/

// @lc code=start
func numIslands(grid [][]byte) int {
	result := 0
	// 非空判断
	if len(grid) == 0 {
		return result
	}

	m, n := len(grid), len(grid[0])

	var changeIslandValue func(i, j int)

	changeIslandValue = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		if grid[i][j] != '1' {
			return
		}

		grid[i][j] = '2'

		changeIslandValue(i-1, j)
		changeIslandValue(i+1, j)
		changeIslandValue(i, j+1)
		changeIslandValue(i, j-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				result++
				changeIslandValue(i, j)
			}
		}
	}

	return result
}

// @lc code=end
