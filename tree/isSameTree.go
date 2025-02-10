package tree

// 判断 2 颗树是否是完全相等的
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

// 判断树是否镜像

// 解法一 dfs
func isSymmetric(root *TreeNode) bool {
	return root == nil || dfs(root.Left, root.Right)
}
func dfs(rootLeft, rootRight *TreeNode) bool {
	if rootLeft == nil && rootRight == nil {
		return true
	}
	if rootLeft == nil || rootRight == nil {
		return false
	}
	if rootLeft.Val != rootRight.Val {
		return false
	}
	return dfs(rootLeft.Left, rootRight.Right) && dfs(rootLeft.Right, rootRight.Left)
}
