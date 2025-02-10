package tree

import "math"

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。假设一个二叉搜索树具有如下特征:
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。

func isValidBST(root *TreeNode) bool {
	// 分别表示负无穷和正无穷
	return isValidbst(root, math.Inf(-1), math.Inf(1))
}

func isValidbst(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	// 左边的所有值，都比当前节点小，右边的所有值，都比当前节点大
	return v < max && v > min &&
		isValidbst(root.Left, min, v) &&
		isValidbst(root.Right, v, max)
}
