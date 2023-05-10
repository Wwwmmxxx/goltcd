package goltcd

import (
	"math"
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 *
 * https://leetcode.cn/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (49.31%)
 * Likes:    1026
 * Dislikes: 0
 * Total Accepted:    295.2K
 * Total Submissions: 598.7K
 * Testcase Example:  '[1,2,3,1]'
 *
 * 峰值元素是指其值严格大于左右相邻值的元素。
 *
 * 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
 *
 * 你可以假设 nums[-1] = nums[n] = -∞ 。
 *
 * 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,1]
 * 输出：2
 * 解释：3 是峰值元素，你的函数应该返回其索引 2。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,1,3,5,6,4]
 * 输出：1 或 5
 * 解释：你的函数可以返回索引 1，其峰值元素为 2；
 * 或者返回索引 5， 其峰值元素为 6。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * -2^31 <= nums[i] <= 2^31 - 1
 * 对于所有有效的 i 都有 nums[i] != nums[i + 1]
 *
 *
 */

// @lc code=start
func findPeakElement(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	for index := range nums {

		if index == 0 {
			if nums[index] > nums[index+1] {
				return index
			}
			continue
		}

		if index == len(nums)-1 {
			if nums[index] > nums[index-1] {
				return index
			}
			continue
		}

		if nums[index] > nums[index-1] && nums[index] > nums[index+1] {
			return index
		}
	}

	return -1
}

func findPeakElement1(nums []int) int {
	f := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}
	left, right := 0, len(nums)

	for {
		mid := (left + right) / 2
		if f(mid) > f(mid-1) && f(mid) > f(mid+1) {
			return mid
		} else if f(mid) < f(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func findPeakElement2(nums []int) int {
	n := len(nums)
	idx := rand.Intn(n)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx) < get(idx+1) {
			idx++
		} else {
			idx--
		}
	}

	return idx
}

// @lc code=end
