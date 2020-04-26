package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) != len(inorder) {
		return nil
	}

	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootIdx := findIndx(inorder, rootVal)

	leftPreorder := []int(nil)
	leftInorder := []int(nil)
	if rootIdx != 0 {
		leftPreorder = preorder[1 : rootIdx+1]
		leftInorder = inorder[:rootIdx]
	}

	rightPreorder := []int(nil)
	rightInorder := []int(nil)
	if len(inorder) > rootIdx+1 {
		rightPreorder = preorder[rootIdx+1:]
		rightInorder = inorder[rootIdx+1:]
	}

	leftNode := buildTree(leftPreorder, leftInorder)
	rightNode := buildTree(rightPreorder, rightInorder)

	return &TreeNode{
		Val:   rootVal,
		Left:  leftNode,
		Right: rightNode,
	}
}

func findIndx(nums []int, target int) int {

	for i, num := range nums {
		if num == target {
			return i
		}
	}

	return -1
}
