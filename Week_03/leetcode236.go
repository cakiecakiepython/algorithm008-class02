package main

type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	leftAncestor := lowestCommonAncestor(root.Left, p, q)
	rightAncestor := lowestCommonAncestor(root.Right, p, q)

	if leftAncestor == p && rightAncestor == q || leftAncestor == q && rightAncestor == p {
		return root
	}

	if leftAncestor == nil {
		return rightAncestor
	}

	return leftAncestor
}
