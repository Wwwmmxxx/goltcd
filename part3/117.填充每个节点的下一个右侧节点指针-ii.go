package goltcd

/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 *
 * https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description/
 *
 * algorithms
 * Medium (65.62%)
 * Likes:    714
 * Dislikes: 0
 * Total Accepted:    184.5K
 * Total Submissions: 281.1K
 * Testcase Example:  '[1,2,3,4,5,null,7]'
 *
 * 给定一个二叉树：
 *
 *
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 *
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
 *
 * 初始状态下，所有 next 指针都被设置为 NULL 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,4,5,null,7]
 * 输出：[1,#,2,3,#,4,5,7,#]
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next
 * 指针连接），'#' 表示每层的末尾。
 *
 * 示例 2：
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
 * 树中的节点数在范围 [0, 6000] 内
 * -100 <= Node.val <= 100
 *
 *
 * 进阶：
 *
 *
 * 你只能使用常量级额外空间。
 * 使用递归解题也符合要求，本题中递归程序的隐式栈空间不计入额外空间复杂度。
 *
 *
 *
 *
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	var f func(...*Node)
	f = func(ns ...*Node) {
		if len(ns) == 0 {
			return
		}

		nextNodes := make([]*Node, 0, len(ns)*2)
		for i := 0; i < len(ns); i++ {
			if i != len(ns)-1 {
				ns[i].Next = ns[i+1]
			}
			if ns[i].Left != nil {
				nextNodes = append(nextNodes, ns[i].Left)
			}
			if ns[i].Right != nil {
				nextNodes = append(nextNodes, ns[i].Right)
			}
		}

		f(nextNodes...)
	}
	f(root)
	return root
}

// @lc code=end
