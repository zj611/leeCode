package tree

// 获取树的最大深度

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 获取树最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// 判断树是否镜像
func isMirror(root *TreeNode) bool {
	return root == nil || dfs1(root.Left, root.Right)
}

func dfs1(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		return p.Val == q.Val && dfs1(p.Left, q.Right) && dfs1(p.Right, q.Left)
	} else {
		return false
	}

}
