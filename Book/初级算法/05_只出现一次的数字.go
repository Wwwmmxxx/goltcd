package main

import "sort"

/**

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明:
	你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
	输入: [2,2,1]
	输出: 1

示例 2:
	输入: [4,1,2,1,2]
	输出: 4

*/

func singleNumber1(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		if i == 0 && nums[i] != nums[i+1] {
			return nums[i]
		}

		if i == 0 {
			continue
		}

		if i == len(nums)-1 && nums[i-1] != nums[i] {
			return nums[i]
		}

		if nums[i-1] == nums[i] || nums[i+1] == nums[i] {
			continue
		}

		return nums[i]

	}

	return 0
}

/**
答案是使用位运算. 对于这道题, 可使用异或运算⊕. 异或运算有以下三个性质:
	1. 任何数和 0 做异或运算, 结果仍然是原来的数, 即 a⊕0=a.
	2. 任何数和其自身做异或运算, 结果是 0, 即 a \oplus a⊕a=0.
	3. 异或运算满足交换律和结合律, 即 a⊕b⊕a = b⊕a⊕a = b⊕(a⊕a)= b⊕0=b.
*/

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
