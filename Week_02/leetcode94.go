package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Stack struct {
	buf []*TreeNode
}

func (s *Stack) Peek() *TreeNode {
	if len(s.buf) == 0 {
		return nil
	}

	return s.buf[len(s.buf)-1]
}

func (s *Stack) Push(node *TreeNode) {
	s.buf = append(s.buf, node)
}

func (s *Stack) Pop() *TreeNode {
	if len(s.buf) == 0 {
		return nil
	}

	res := s.buf[len(s.buf)-1]
	s.buf = s.buf[:len(s.buf)-1]
	return res
}

func inorderTraversal(root *TreeNode) []int {

	res := []int{}
	stack := Stack{}
	curr := root

	for curr != nil || stack.Peek() != nil {

		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}

		curr = stack.Pop()
		res = append(res, curr.Val)
		curr = curr.Right
	}

	return res
}
