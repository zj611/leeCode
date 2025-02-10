package tree

// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
// 输出 1~n 元素组成的 BST 所有解。这一题递归求解即可。外层循环遍历 1~n 所有结点，作为根结点，内层双 层递归分别求出左子树和右子树。

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateBSTree(1, n)
}

func generateBSTree(start, end int) []*TreeNode {
	var tree []*TreeNode
	if start > end {
		tree = append(tree, nil)
		return tree
	}
	for i := start; i <= end; i++ {
		//递归生成以 i 为根节点的左右子树
		left := generateBSTree(start, i-1)
		right := generateBSTree(i+1, end)
		// 遍历 left 和 right 中的每个子树组合：
		// 创建新的根节点 root，值为 i，左子树为 l，右子树为 r。
		// 将生成的树 root 添加到结果切片 tree 中。
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				tree = append(tree, root)
			}
		}
	}
	return tree
}
