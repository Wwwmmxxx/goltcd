package goltcd

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
 *
 * https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (57.52%)
 * Likes:    763
 * Dislikes: 0
 * Total Accepted:    302.3K
 * Total Submissions: 525.6K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[3],[20,9],[15,7]]
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
 * -100 <= Node.val <= 100
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var (
		res = [][]int{}
		f   func(int, ...*TreeNode)
	)

	f = func(dept int, nodes ...*TreeNode) {
		if len(nodes) == 0 {
			return
		}

		r := make([]int, 0, len(nodes))
		for _, n := range nodes {
			r = append(r, n.Val)
		}
		res = append(res, r)

		// 开始考虑下一个节点该怎么做?
		tmpNodes := make([]*TreeNode, 0, len(nodes)*2)
		for i := len(nodes) - 1; i >= 0; i-- {
			if dept%2 == 0 {
				if nodes[i].Right != nil {
					tmpNodes = append(tmpNodes, nodes[i].Right)
				}
				if nodes[i].Left != nil {
					tmpNodes = append(tmpNodes, nodes[i].Left)
				}
			} else {
				if nodes[i].Left != nil {
					tmpNodes = append(tmpNodes, nodes[i].Left)
				}
				if nodes[i].Right != nil {
					tmpNodes = append(tmpNodes, nodes[i].Right)
				}
			}
		}

		f(dept+1, tmpNodes...)
	}

	f(0, root)

	return res
}

// func zigzagLevelOrder(root *TreeNode) [][]int {
//     res := [][]int{}

//     var dfs func(*TreeNode, int)

//     dfs = func(node *TreeNode, depth int){
//         if node ==nil{
//             return
//         }
//         if len(res)<=depth{
//             res = append(res, []int{})
//         }

//         if depth&1==0{
//             res[depth] = append(res[depth], node.Val)
//         } else{
//             res[depth] = append([]int{node.Val}, res[depth]...)
//         }

//         dfs(node.Left, depth+1)
//         dfs(node.Right, depth+1)

//     }

//     dfs(root, 0)
//     return res
// }

// @lc code=end
