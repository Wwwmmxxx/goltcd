package goltcd

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode.cn/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (52.48%)
 * Likes:    1841
 * Dislikes: 0
 * Total Accepted:    885K
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
 *
 * 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
 *
 * 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m
 * 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
 * 输出：[1,2,2,3,5,6]
 * 解释：需要合并 [1,2,3] 和 [2,5,6] 。
 * 合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1], m = 1, nums2 = [], n = 0
 * 输出：[1]
 * 解释：需要合并 [1] 和 [] 。
 * 合并结果是 [1] 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums1 = [0], m = 0, nums2 = [1], n = 1
 * 输出：[1]
 * 解释：需要合并的数组是 [] 和 [1] 。
 * 合并结果是 [1] 。
 * 注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m + n
 * nums2.length == n
 * 0 <= m, n <= 200
 * 1 <= m + n <= 200
 * -10^9 <= nums1[i], nums2[j] <= 10^9
 *
 *
 *
 *
 * 进阶：你可以设计实现一个时间复杂度为 O(m + n) 的算法解决此问题吗？
 *
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 && n == 1 {
		nums1[0] = nums2[0]
		return
	}

	moveForward := func(nums []int, index int) []int {
		for i := len(nums) - 2; i >= index; i-- {
			nums[i+1] = nums[i]
		}
		return nums
	}
	n1Index, n2Index := 0, 0
	for ; n1Index < len(nums1); n1Index++ {
		// 如果最后是nums2先遍历完了, 则直接返回, 因为此时num1和num2已经成功合并
		if n2Index == len(nums2) {
			return
		}

		if nums2[n2Index] < nums1[n1Index] {
			moveForward(nums1, n1Index)
			nums1[n1Index] = nums2[n2Index]
			n2Index++
		}
	}

	// 可能会有第二种情况, 就是nums2没有遍历完, 如果没有遍历完, 就手动附上
	for i, j := len(nums2)-1, len(nums1)-1; i >= n2Index; {
		nums1[j] = nums2[i]
		i--
		j--
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	for k := m + n - 1; k >= m; k-- {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[k] = nums1[m-1]
			m--
		} else {
			nums1[k] = nums2[n-1]
			n--
		}
	}
}

// @lc code=end
