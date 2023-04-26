package goltcd

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode.cn/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (71.28%)
 * Likes:    1804
 * Dislikes: 0
 * Total Accepted:    615.9K
 * Total Submissions: 864K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 100] 内
 * 0 <= Node.val <= 100
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		realH    = &ListNode{Next: head}
		pre, cur = realH, realH.Next
	)

	for cur != nil {
		// 如果cur.Next不为空, 则和下一个进行交换
		if cur.Next != nil {
			next2 := cur.Next.Next
			pre.Next = cur.Next
			pre.Next.Next = cur
			cur.Next = next2
		}
		pre = cur
		cur = cur.Next
	}

	return realH.Next
}

// @lc code=end
