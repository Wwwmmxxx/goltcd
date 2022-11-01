package main

/**

给定一个整数数组nums和一个整数目标值target, 请你在该数组中找出"和"为目标值target的那两个整数, 并返回它们的数组下标.
你可以假设每种输入只会对应一个答案. 但是, 数组中同一个元素在答案里不能重复出现.
你可以按任意顺序返回答案。

示例 1:
	输入: nums = [2,7,11,15], target = 9
	输出: [0,1]
	解释: 因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2:
	输入: nums = [3,2,4], target = 6
	输出: [1,2]

示例 3:
	输入: nums = [3,3], target = 6
	输出: [0,1]

提示：
	2 <= nums.length <= 104
	-109 <= nums[i] <= 109
	-109 <= target <= 109
	只会存在一个有效答案

进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2jrse/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func twoSum(nums []int, target int) []int {

	// key: 值, value: 下表
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		n2 := target - n1
		if numsIndex, isExist := m[n2]; isExist {
			return []int{numsIndex, i}
		} else {
			m[n1] = i
		}
	}
	return nil
}
