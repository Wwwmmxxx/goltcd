package goltcd

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (72.17%)
 * Likes:    1018
 * Dislikes: 0
 * Total Accepted:    282.1K
 * Total Submissions: 391K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder
 * 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
 * 输出：[3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入：inorder = [-1], postorder = [-1]
 * 输出：[-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= inorder.length <= 3000
 * postorder.length == inorder.length
 * -3000 <= inorder[i], postorder[i] <= 3000
 * inorder 和 postorder 都由 不同 的值组成
 * postorder 中每一个值都在 inorder 中
 * inorder 保证是树的中序遍历
 * postorder 保证是树的后序遍历
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
//  inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// func buildTree(inorder []int, postorder []int) *TreeNode {
// 	var f func(inorder, postorder []int) *TreeNode
// 	f = func(inorder, postorder []int) *TreeNode {
// 		node := &TreeNode{Val: postorder[len(postorder)-1]}

// 		// 找到
// 		var index int
// 		for i := range inorder {
// 			if inorder[i] == node.Val {
// 				index = i
// 				break
// 			}
// 		}

// 		// left
// 		if index-1 >= 0 {
// 			node.Left = f(inorder[:index-1], postorder[:len(postorder)-1])
// 		}
// 		if index+1 < len(inorder) {
// 			node.Right = f(inorder[index+1:], postorder[:len(postorder)-1])
// 		}

// 		return node
// 	}

// 	return f(inorder, postorder)
// }

func buildTree(inorder []int, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}

		// 后序遍历的末尾元素即为当前子树的根节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		inorderRootIndex := idxMap[val]
		root.Right = build(inorderRootIndex+1, inorderRight)
		root.Left = build(inorderLeft, inorderRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}

// @lc code=end
