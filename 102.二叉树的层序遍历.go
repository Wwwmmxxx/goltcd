package goltcd

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (65.52%)
 * Likes:    1643
 * Dislikes: 0
 * Total Accepted:    787.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[3],[9,20],[15,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：[[1]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 2000] 内
 * -1000 <= Node.val <= 1000
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	result := [][]int{
		{root.Val},
	}

	var f func(dept int, nodes ...*TreeNode)

	f = func(dept int, nodes ...*TreeNode) {
		if len(nodes) == 0 {
			return
		}

		var tmpNodes = make([]*TreeNode, 0)
		var curDeptValue = make([]int, 0, len(nodes))
		for i := 0; i < len(nodes); i++ {

			if nodes[i].Left != nil {
				curDeptValue = append(curDeptValue, nodes[i].Left.Val)
				tmpNodes = append(tmpNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				curDeptValue = append(curDeptValue, nodes[i].Right.Val)
				tmpNodes = append(tmpNodes, nodes[i].Right)
			}

		}

		if len(curDeptValue) > 0 {
			result = append(result, curDeptValue)
		}

		f(dept+1, tmpNodes...)
	}

	f(0, root)

	return result
}

// @lc code=end
