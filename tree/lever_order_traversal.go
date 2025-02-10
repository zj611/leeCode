package tree

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	p := []*TreeNode{root}
	for i := 0; len(p) > 0; i++ {
		// 开始第一层，新加一个数组（每层用一个数组记录该层的node值）
		ret = append(ret, []int{})
		var q []*TreeNode
		for j := 0; j < len(p); j++ {
			node := p[j]
			ret[i] = append(ret[i], node.Val)
			// 记录完后，把该层的node都追加到q数组内，用于下一个循环
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		// p遍历完毕（该层结束），把该层的q数组赋给p，重新开始一轮循环
		p = q
	}
	return ret
}
