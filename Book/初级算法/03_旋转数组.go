package book01

/**

给你一个数组, 将数组中的元素向右轮转"k"个位置, 其中"k"是非负数.

示例 1:
	输入: nums = [1,2,3,4,5,6,7], k = 3
	输出: [5,6,7,1,2,3,4]
	解释:
		向右轮转 1 步: [7,1,2,3,4,5,6]
		向右轮转 2 步: [6,7,1,2,3,4,5]
		向右轮转 3 步: [5,6,7,1,2,3,4]

示例2:
	输入：nums = [-1,-100,3,99], k = 2
	输出：[3,99,-1,-100]
	解释:
		向右轮转 1 步: [99,-1,-100,3]
		向右轮转 2 步: [3,99,-1,-100]

提示：
	1 <= nums.length <= 105
	-231 <= nums[i] <= 231 - 1
	0 <= k <= 105

进阶：
	尽可能想出更多的解决方案, 至少有"三种"不同的方法可以解决这个问题。
	你可以使用空间复杂度为O(1)的原地算法解决这个问题吗？

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2skh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func rotate1(nums []int, k int) {

	length := len(nums)

	tempNums := make([]int, len(nums))

	copy(tempNums, nums)

	for i := 0; i < len(nums); i++ {
		nums[(i+k-1)%length] = tempNums[i]
	}

}

/**
数组反转做法
*/
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
