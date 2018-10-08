package main

import(
	"fmt"
)


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
    var rst []int
	h := 0
	if root == nil{
		return nil
	}
	rst = preSearch(root, rst, h)
	return rst

}

func preSearch(root *TreeNode, rst []int, h int) []int{
	fmt.Println(rst, h)
	if len(rst) <= h {
		rst = append(rst, root.Val)
	}else{
		rst[h] = root.Val
	}

	if root.Left != nil{
        rst = preSearch(root.Left, rst, h+1)
    }

	if root.Right != nil{
        rst = preSearch(root.Right, rst, h+1)
    }
	return rst

}


func newTree(in []interface{}, rr *TreeNode) *TreeNode{
	var list []*TreeNode
	index := 0
	l := len(in)
	if l == 0 {
		return nil
	}
	root := &TreeNode{
			Val : in[index].(int),
	}
	list = append(list, root)
	index++
	for {
		if index >= l{
			break
		}
		tmp := list[0]
		var L, R *TreeNode 
		if index < l && in[index] != nil{
			L = &TreeNode{
				Val : in[index].(int),
			}
			list = append(list, L)
		}
		index++
		if index < l && in[index] != nil{
            R = &TreeNode{
                Val : in[index].(int),
        	}
			list = append(list, R)
		}
		index++

		tmp.Left = L
		tmp.Right = R
		list = list[1:]
		//fmt.Println(tmp, tmp.Left, tmp.Right)
	}
	return root 


}

func main(){
	//in := []interface{}{1,nil,4,3,5,2,nil,nil,6}
	//in := []interface{}{1,2,3,nil,5,nil,4}
	in := []interface{}{1, 2}
	root := newTree(in, nil)
	fmt.Println(rightSideView(root))
}
