package tree

import "math"

// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，
// 这条路径上所有节点值相加等于目标和。说明: 叶子节点是指没有子节点的节点。

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

// 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
// 说明: 叶子节点是 指没有子节点的节点

func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	slice = findPath(root, sum, slice, []int(nil))
	return slice
}

func findPath(n *TreeNode, sum int, slice [][]int, stack []int) [][]int {
	if n == nil {
		return slice
	}
	sum -= n.Val
	stack = append(stack, n.Val)
	if sum == 0 && n.Left == nil && n.Right == nil {
		slice = append(slice, append([]int{}, stack...))
		// 回溯，将当前节点从路径中移除，拿掉最后一个，注意是前闭后开原则
		stack = stack[:len(stack)-1]
	}
	slice = findPath(n.Left, sum, slice, stack)
	slice = findPath(n.Right, sum, slice, stack)
	return slice
}

func pathSum1(root *TreeNode, sum int) [][]int {
	var slice [][]int
	var stack []int
	return findpath1(root, sum, slice, stack)
}

func findpath1(node *TreeNode, sum int, slice [][]int, stack []int) [][]int {

	if node == nil {
		return slice
	}
	sum -= node.Val
	stack = append(stack, node.Val)

	if sum == 0 && node.Left == nil && node.Right == nil {
		slice = append(slice, append([]int{}, stack...))
		stack = stack[:len(stack)-1]
	}
	slice = findpath1(node.Left, sum, slice, stack)
	slice = findpath1(node.Right, sum, slice, stack)
	return slice
}

// 给定一个非空二叉树，返回其最大路径和。本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。
// 该路径至少包含一个节点，且不一定经过根节点。

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxValue := math.MinInt32
	getPathSum(root, &maxValue)
	return maxValue
}
func getPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return math.MinInt32
	}
	left := getPathSum(root.Left, maxSum)
	right := getPathSum(root.Right, maxSum)

	//currMax 表示以当前节点为起点的路径的最大值，有三种情况：
	//•	包括左子树路径和 (left+root.Val)。
	//•	包括右子树路径和 (right+root.Val)。
	//•	仅包括当前节点的值 (root.Val)。其实还有左+右+root.Val
	currMax := max(max(left+root.Val, right+root.Val), root.Val)
	//maxSum 被更新为全局最大值，考虑三种情况：
	//•	当前节点的路径和（currMax）。
	//•	左右子树加当前节点形成的路径（left+right+root.Val）。
	//5.	返回值：
	*maxSum = max(*maxSum, max(currMax, left+right+root.Val))
	return currMax
}

func maxPathSum1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxValue := math.MinInt32
	getPathSum1(root, &maxValue)
	return maxValue
}

func getPathSum1(root *TreeNode, maxSum *int) int {
	if root == nil {
		return math.MinInt32
	}
	leftVal := getPathSum1(root.Left, maxSum)
	rightVal := getPathSum1(root.Right, maxSum)

	curMax := max(max(leftVal+root.Val, rightVal+root.Val), root.Val)
	return max(*maxSum, max(curMax, leftVal+rightVal+root.Val))
}

//给定一个二叉树，原地将它展开为链表。
//要求把二叉树“打平”，按照先根遍历的顺序，把树的结点都放在右结点中。
//按照递归和非递归思路实现即可。
//递归的思路可以这么想:倒序遍历一颗树，即是先遍历右孩子，然后遍历左孩子，最后再遍历根节点。

func flatten1(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	flatten1(root.Left)
	flatten1(root.Right)
	//•	保存当前右子树：
	//用变量 currRight 保存当前节点的右子树，防止被后续操作覆盖。
	currRight := root.Right
	//•	移动左子树：
	//将已经被扁平化的左子树移动到当前节点的右子树位置。
	root.Right = root.Left
	//•	清空左子树：
	//将当前节点的左子树置为 nil，保证最终结果是一个右倾链表。
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = currRight
}
