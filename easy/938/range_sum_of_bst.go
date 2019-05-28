package rangeSumOfBST

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var total int
	if root == nil {
		return total
	}
	add(root, &L, &R, &total)
	return total

}

func add(treeNode *TreeNode, L *int, R *int, total *int) {
	if treeNode == nil {
		return
	}

	if treeNode.Val >= *L && treeNode.Val <= *R {
		*total += treeNode.Val
	}
	if treeNode.Left != nil {
		add(treeNode.Left, L, R, total)
	}
	if treeNode.Right != nil {
		add(treeNode.Right, L, R, total)
	}

}
