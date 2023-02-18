package book01

import "sort"

/**

给你两个整数数组nums1和nums2, 请你以数组形式返回两数组的交集.
返回结果中每个元素出现的次数, 应与元素在两个数组中都出现的次数一致(如果出现次数不一致，则考虑取较小值).可以不考虑输出结果的顺序.

示例 1:
	输入: nums1 = [1,2,2,1], nums2 = [2,2]
	输出: [2,2]

示例 2:
	输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	输出: [4,9]

提示:
	1 <= nums1.length, nums2.length <= 1000
	0 <= nums1[i], nums2[i] <= 1000

进阶:
	1. 如果给定的数组已经排好序呢? 你将如何优化你的算法?
	2. 如果nums1的大小比nums2小, 哪种方法更优？
	3. 如果nums2的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2skh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func intersect(nums1 []int, nums2 []int) (output []int) {

	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {

		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			output = append(output, nums1[i])
			i++
			j++
		}

	}

	return
}
