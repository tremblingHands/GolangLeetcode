/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
	
	//fmt.Println(rst, h)
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
