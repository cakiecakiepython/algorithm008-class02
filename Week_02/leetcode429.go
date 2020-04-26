package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {

	if root == nil {
		return nil
	}

	helper := []*Node{root}
	res := [][]int{}

	for len(helper) != 0 {

		resElems := []int{}
		tmp := []*Node{}
		for _, node := range helper {

			tmp = append(tmp, node.Children...)
			resElems = append(resElems, node.Val)
		}

		res = append(res, resElems)
		helper = tmp
	}

	return res
}
