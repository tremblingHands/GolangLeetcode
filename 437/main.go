package main

import(
	"fmt"
)


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil{
		return 0
	}
	return cnt(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func cnt(root *TreeNode, sum int) int{
	if root == nil {
		return 0
	}
	ret := cnt(root.Left, sum - root.Val) + cnt(root.Right, sum - root.Val)
	if root.Val == sum {
		return 1 + ret 
	}else {
		return ret 
	}


}

