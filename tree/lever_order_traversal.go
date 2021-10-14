package tree

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	p := []*TreeNode{root}
	for i := 0; len(p) > 0; i++ {
		ret = append(ret, []int{})
		q := []*TreeNode{}
		for j := 0; j < len(p); j++ {
			node := p[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		p = q
	}
	return ret
}
