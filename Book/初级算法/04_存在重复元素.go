package main

import "sort"

/**
给你一个整数数组nums. 如果任一值在数组中出现"至少两次", 返回true; 如果数组中每个元素互不相同, 返回 false.

示例 1:
	输入: nums = [1,2,3,1]
	输出: true

示例 2:
	输入: nums = [1,2,3,4]
	输出: false

示例 3:
	输入: nums = [1,1,1,3,3,4,3,2,4,2]
	输出: true

*/

func containsDuplicate1(nums []int) bool {

	hasValidate := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {

		isExist := hasValidate[nums[i]]
		if !isExist {
			for j := i + 1; j < len(nums); j++ {
				if nums[i] == nums[j] {
					return true
				}
			}
		}
		hasValidate[nums[i]] = true

	}

	return false
}

func containsDuplicate2(nums []int) bool {

	hasValidate := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {

		isExist := hasValidate[nums[i]]
		if !isExist {
			hasValidate[nums[i]] = true
		} else {
			return true
		}

	}

	return false
}

func containsDuplicate(nums []int) bool {

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {

		if nums[i] == nums[i-1] {
			return true
		}

	}

	return false
}
