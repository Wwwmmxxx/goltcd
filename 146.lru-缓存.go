package goLeetCode

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 *
 * https://leetcode.cn/problems/lru-cache/description/
 *
 * algorithms
 * Medium (53.44%)
 * Likes:    2636
 * Dislikes: 0
 * Total Accepted:    469.7K
 * Total Submissions: 878.8K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
 *
 * 实现 LRUCache 类：
 *
 *
 *
 *
 * LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 * void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
 * key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
 *
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * 输出
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * 解释
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // 缓存是 {1=1}
 * lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
 * lRUCache.get(1);    // 返回 1
 * lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
 * lRUCache.get(2);    // 返回 -1 (未找到)
 * lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
 * lRUCache.get(1);    // 返回 -1 (未找到)
 * lRUCache.get(3);    // 返回 3
 * lRUCache.get(4);    // 返回 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= capacity <= 3000
 * 0 <= key <= 10000
 * 0 <= value <= 10^5
 * 最多调用 2 * 10^5 次 get 和 put
 *
 *
*/

// @lc code=start
type LRUCache struct {
	capacity   int
	m          map[int]*LRUNode
	Head, Tail *LRUNode
}

type LRUNode struct {
	Key       int
	Value     int
	Pre, Next *LRUNode
}

func Constructor(capacity int) LRUCache {

	head, tail := &LRUNode{}, &LRUNode{}

	head.Next = tail
	tail.Pre = head

	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*LRUNode),
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) deleteNode(node *LRUNode) {
	node.Next.Pre = node.Pre
	node.Pre.Next = node.Next
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	// 先删除当前节点
	this.deleteNode(node)
	// 把当前节点放到header下
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *LRUNode) {
	// 处理反向链表
	this.Head.Next.Pre = node
	node.Pre = this.Head
	// 处理正向链表
	node.Next = this.Head.Next
	this.Head.Next = node
}

func (this *LRUCache) removeTail() *LRUNode {
	removeNode := this.Tail.Pre
	this.deleteNode(removeNode)
	return removeNode
}

func (this *LRUCache) Get(key int) int {
	if node, isExist := this.m[key]; isExist {
		this.moveToHead(node)
		return node.Value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if node, isExist := this.m[key]; isExist {
		node.Value = value
		this.moveToHead(node)
		return
	}

	// 如果已经满了, 先删除header
	if this.capacity == len(this.m) {
		removeNode := this.removeTail()
		delete(this.m, removeNode.Key)
	}

	newNode := &LRUNode{
		Key:   key,
		Value: value,
	}

	this.addToHead(newNode)
	this.m[key] = newNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
