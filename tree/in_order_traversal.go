package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//深度优先遍历之中序遍历
//var res []int
//func inorderTraversal(root *TreeNode) []int {
//	res = make([]int,0)
//	inOrder(root)
//	return res
//}
//
//func inOrder(root *TreeNode){
//	if root == nil{
//		return
//	}
//	inOrder(root.Left)
//	res = append(res,root.Val)
//	inOrder(root.Right)
//	return
//}

func inorderTraversal(root *TreeNode) (res []int) {
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}

	inOrder(root)
	return
}
