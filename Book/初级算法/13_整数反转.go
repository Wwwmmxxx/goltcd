package main

import (
	"math"
	"strconv"
	"strings"
)

/**

给你一个 32 位的有符号整数x, 返回将 x 中的数字部分反转后的结果.
如果反转后整数超过 32 位的有符号整数的范围[−2^31, 2^31− 1], 就返回 0.
假设环境不允许存储 64 位整数(有符号或无符号)

示例 1:
	输入: x = 123
	输出: 321

示例 2:
	输入: x = -123
	输出: -321

示例 3：
	输入: x = 120
	输出: 21

示例 4：
	输入: x = 0
	输出: 0

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnx13t/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func reverse2(x int) int {
	ret := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		ret = ret*10 + pop
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}
	return ret
}

func reverse2_1(x int) (res int) {

	defer func() {
		if (x < math.MinInt32 || x > math.MaxInt32) ||
			(res < math.MinInt32 || res > math.MaxInt32) {
			res = 0
		}
	}()

	if x == 0 {
		return 0
	}

	bytes := []byte(strconv.FormatInt(int64(x), 10))
	var strBuilder strings.Builder

	if bytes[0] == '-' {
		strBuilder.WriteByte('-')
		bytes = bytes[1:]
	}

	for i := len(bytes); i > 0; i-- {
		strBuilder.WriteByte(bytes[i-1])
	}

	revNum, _ := strconv.Atoi(strBuilder.String())

	return revNum
}
