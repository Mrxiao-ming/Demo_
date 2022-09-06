package study

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 21. 合并两个有序链表
func Fun21(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		pre = &ListNode{
			Val:  0,
			Next: nil,
		}
		res = pre
	)
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 == nil {
		pre.Next = list2
	}
	if list2 == nil {
		pre.Next = list1
	}
	return res.Next
}

// 70. 爬楼梯
func Fun70(n int) int {
	if n <= 2 {
		return n
	}
	var (
		one = 1
		tow = 2
	)
	for i := 3; i < n; i++ {
		one, tow = tow, one+tow
	}
	return tow + one
}

// 98. 验证二叉搜索树
func Fun98_1(root *TreeNode) bool { // 递归
	return Fun98Recursive(root, math.MinInt64, math.MaxInt64)
}
func Fun98Recursive(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return Fun98Recursive(root.Left, lower, root.Val) && Fun98Recursive(root.Right, root.Val, upper)
}
func Fun98_2(root *TreeNode) bool { // 中序法
	var (
		min   = math.MinInt64
		stack []*TreeNode
	)
	for len(stack) > 0 || root != nil {
		// 将左节点全部压入栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 比较
		if min >= root.Val {
			return false
		}
		// min & root
		min = root.Val
		root = root.Right
	}
	return true
}

// 121. 买卖股票的最佳时机
func Fun121_1(prices []int) int {
	var maxPrice int
	for index, price := range prices {
		for i := index + 1; i < len(prices); i++ {
			if prices[i]-price > maxPrice {
				maxPrice = prices[i] - price
			}
		}
	}
	return maxPrice
}
func Fun121_2(prices []int) int {
	var (
		minPrice     = math.MaxInt64
		maxPriceDiff int
	)
	for _, price := range prices {
		// 找到区间范围内最大的值和最小的值
		if price < minPrice {
			minPrice = price
		}
		// 找最大值的同时要使最大值在最小值后
		if price-minPrice > maxPriceDiff {
			maxPriceDiff = price - minPrice
		}
	}
	return maxPriceDiff
}

// 102. 二叉树的层序遍历
func Fun102(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		res      [][]int
		nodeList = []*TreeNode{root}
	)
	for len(nodeList) > 0 {
		var (
			nodes   []*TreeNode
			valList []int
		)
		for _, node := range nodeList {
			valList = append(valList, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		res = append(res, valList)
		nodeList = nodes
	}
	return res
}

// 142. 环形链表 II
func Fun142_1(head *ListNode) *ListNode {
	var (
		distMap = make(map[*ListNode]struct{})
	)
	for head != nil {
		if _, ok := distMap[head]; ok {
			return head
		}
		distMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}
func Fun142_2(head *ListNode) *ListNode {
	var fast, slow = head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// 200. 岛屿数量
func Fun200(grid [][]byte) int {
	var res int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if string(grid[y][x]) == "1" {
				res++
				Fun200Dfs(grid, x, y)
			}
		}
	}
	return res
}
func Fun200Dfs(grid [][]byte, x, y int) {
	grid[y][x] = 0
	if x-1 >= 0 && string(grid[y][x-1]) == "1" {
		Fun200Dfs(grid, x-1, y)
	}
	if x+1 < len(grid[0]) && string(grid[y][x+1]) == "1" {
		Fun200Dfs(grid, x+1, y)
	}
	if y-1 >= 0 && string(grid[y-1][x]) == "1" {
		Fun200Dfs(grid, x, y-1)
	}
	if y+1 < len(grid) && string(grid[y+1][x]) == "1" {
		Fun200Dfs(grid, x, y+1)
	}
}

// 205. 同构字符串
func Fun205(s string, t string) bool {
	var (
		sMap = make(map[byte]byte)
		tMap = make(map[byte]byte)
	)
	if len(s) == 0 || len(t) == 0 || len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		tStr, sStr := sMap[s[i]], tMap[t[i]]
		if (sStr > 0 && sStr != s[i]) || (tStr > 0 && tStr != t[i]) {
			return false
		}

		sMap[s[i]] = t[i]
		tMap[t[i]] = s[i]

	}
	return true
}

// 206
func Fun206(head *ListNode) *ListNode {
	var (
		pre = &ListNode{
			Val:  head.Val,
			Next: nil,
		}
		next = head.Next
	)
	for next != nil {
		next_ := next.Next
		next.Next = pre
		pre, next = next, next_

	}
	return pre
}

// 235. 二叉搜索树的最近公共祖先
func Fun235(root, p, q *TreeNode) *TreeNode {
	var (
		rootQ  = root
		rootP  = root
		qStack []*TreeNode
		pMap   = make(map[*TreeNode]struct{})
	)
	// q
	for rootQ != nil {
		qStack = append(qStack, rootQ)

		if rootQ.Val == q.Val {
			break
		}

		if q.Val < rootQ.Val {
			rootQ = rootQ.Left
		} else {
			rootQ = rootQ.Right
		}
	}
	// p
	for rootP != nil {
		pMap[rootP] = struct{}{}

		if rootP.Val == p.Val {
			break
		}

		if p.Val < rootP.Val {
			rootP = rootP.Left
		} else {
			rootP = rootP.Right
		}
	}

	// 弹栈
	for len(qStack) > 0 {
		if _, ok := pMap[qStack[len(qStack)-1]]; ok {
			return qStack[len(qStack)-1]
		}
		qStack = qStack[:len(qStack)-1]
	}

	return nil
}

// 278. 第一个错误的版本
func Func279IsBadVersion(n int) bool {
	return false
}
func Func278(n int) int {
	var (
		left  = 1
		right = n
	)
	for left < right {
		if Func279IsBadVersion(left) {
			return left
		}
		mid := left + (right-left)/2 // 防止计算时溢出
		if Func279IsBadVersion(mid) {
			left++
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 392. 判断子序列
func Fun392(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	var (
		bys        = []byte(s)
		startIndex = 0
		endIndex   = len(bys) - 1
	)

	for _, byt := range t {
		if startIndex <= endIndex {
			if bys[startIndex] == byte(byt) {
				startIndex++
			}
		}
	}

	if startIndex == endIndex+1 {
		return true
	}
	return false
}

// 409. 最长回文串
func Fun409(s string) int {
	var (
		strMap = make(map[byte]int)
		count  int
	)
	for _, byt := range []byte(s) {
		if num, ok := strMap[byt]; ok {
			strMap[byt] = num + 1
		} else {
			strMap[byt] = 1
		}
	}
	for _, val := range strMap {
		count += val - val%2
	}
	if count < len(s) {
		count += 1
	}
	return count
}

// 509. 斐波那契数
func Fun509(n int) int {
	if n < 2 {
		return n
	}
	var (
		first = 0
		sec   = 1
	)
	for i := 2; i < n; i++ {
		first, sec = sec, sec+first
	}
	return first + sec
}

// 589. N 叉树的前序遍历
func Fun589(root *Node) []int {
	if root == nil {
		return nil
	}
	var allNode []int
	allNode = append(allNode, root.Val)
	return Fun589Child(root, allNode)
}
func Fun589Child(root *Node, allNode []int) []int {
	for _, child := range root.Children {
		allNode = append(allNode, child.Val)
		if child.Children != nil {
			allNode = Fun589Child(child, allNode)
		}
	}
	return allNode
}

// 704. 二分查找
func Fun704(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)
	for left <= right {
		mid := (right - left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[mid] < target {
			left = mid + 1
			right--
		} else {
			right = mid - 1
			left++
		}
	}
	return -1
}

// 724. 寻找数组的中心下标
func Fun724(nums []int) int {
	var (
		total int
		sum   int
	)
	for _, val := range nums {
		total += val
	}

	for index, val := range nums {
		if 2*sum+val == total {
			return index
		}
		sum += val
	}

	return -1
}

// 733. 图像渲染
func Fun733(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 {
		return image
	}
	var (
		matchStack = [][]int{{sr, sc}}
		column     = len(image) - 1
		row        = len(image[0]) - 1
		matchColor = image[sr][sc]
	)
	if matchColor == color {
		return image
	}
	for len(matchStack) > 0 {
		// 弹栈
		rc := matchStack[len(matchStack)-1]
		matchStack = matchStack[:len(matchStack)-1]
		sr = rc[0]
		sc = rc[1]
		image[sr][sc] = color
		if sr-1 >= 0 {
			if image[sr-1][sc] == matchColor {
				matchStack = append(matchStack, []int{sr - 1, sc})
			}
		}
		if sr+1 <= column {
			if image[sr+1][sc] == matchColor {
				matchStack = append(matchStack, []int{sr + 1, sc})
			}
		}
		if sc-1 >= 0 {
			if image[sr][sc-1] == matchColor {
				matchStack = append(matchStack, []int{sr, sc - 1})
			}
		}
		if sc+1 <= row {
			if image[sr][sc+1] == matchColor {
				matchStack = append(matchStack, []int{sr, sc + 1})
			}
		}
	}
	return image
}

// 876. 链表的中间结点
func Fun876(head *ListNode) *ListNode {
	var list []*ListNode
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	if len(list) == 0 {
		return nil
	}
	return list[len(list)/2]
}

// 1480 一维数组的动态和
func Fun1480(nums []int) []int {
	var num int
	for index, val := range nums {
		num += val
		nums[index] = num
	}
	return nums
}
