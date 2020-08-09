package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) (ans []*TreeNode) {
	DFSinOrder(&root, to_delete, true, &ans)
	return ans
}

func DFSinOrder(root **TreeNode, to_delete []int, appendToAnswer bool, ans *[]*TreeNode) {

	// guard
	if *root == nil {
		return
	}

	// delete if its in the list
	tmp := *root
	if contains(to_delete, tmp.Val) {
		appendToAnswer = true
		*root = nil
	} else {
		if appendToAnswer {
			*ans = append(*ans, *root)
			appendToAnswer = false //switch off
		}
	}

	DFSinOrder(&tmp.Left, to_delete, appendToAnswer, ans)
	DFSinOrder(&tmp.Right, to_delete, appendToAnswer, ans)
}

func contains(to_delete []int, search int) bool {
	for _, value := range to_delete {
		if value == search {
			return true
		}
	}
	return false
}
