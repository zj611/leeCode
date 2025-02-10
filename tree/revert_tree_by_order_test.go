package tree

import "fmt"

// buildTree 根据前序和中序遍历恢复二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 前序遍历的第一个值是根节点
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的位置
	rootIndex := 0
	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
			break
		}
	}

	// 切分左子树和右子树
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:] // 已经拿掉了rootIndex

	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[1+len(leftInorder):]

	// 递归构建左子树和右子树
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	rootIndex := 0
	for i, v := range inorder {
		if v == root.Val {
			rootIndex = i
		}
	}

	leftInOrder := inorder[:rootIndex]
	rightInOrder := inorder[rootIndex+1:] // 已经拿掉了rootIndex

	leftPreOrder := preorder[1 : len(leftInOrder)+1]
	rightPreOrder := preorder[len(leftInOrder)+1:]

	root.Left = buildTree1(leftPreOrder, leftInOrder)
	root.Right = buildTree1(rightPreOrder, rightInOrder)
	return root
}

// 前序遍历打印二叉树
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 中序遍历打印二叉树
func inorderTraversal1(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal1(root.Left)
	fmt.Print(root.Val, " ")
	inorderTraversal1(root.Right)
}

func main() {
	// 示例前序和中序遍历数组
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	// 恢复二叉树
	root := buildTree(preorder, inorder)

	// 打印验证
	fmt.Print("前序遍历结果: ")
	preorderTraversal(root)
	fmt.Println()
	fmt.Print("中序遍历结果: ")
	inorderTraversal(root)
}
