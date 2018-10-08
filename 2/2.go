/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func fix(l *ListNode) *ListNode {
	head := l
	flag := 0
	for {
		l.Val += flag
		flag = 0
		if l.Val >9{
			l.Val -= 10
			flag = 1
		}
		if l.Next == nil{
			break
		}
		l = l.Next
	}

	if flag == 1{
		tail := &ListNode{
			Val : flag,
			Next : nil,
		}
		l.Next = tail
	}

	return head
}


func add(l1 *ListNode, l2 *ListNode) *ListNode {
        var rst *ListNode = new(ListNode)
        var head *ListNode = rst
        for l1 != nil && l2 != nil {
                rst.Next = &ListNode{Val: l1.Val + l2.Val}
                rst = rst.Next
                l1 = l1.Next
                l2 = l2.Next
        }
        if l1 != nil {
                for l1 != nil{
                        rst.Next, l1 = l1, l1.Next
                        rst = rst.Next
                }
        }else {
                for l2 != nil{
                        rst.Next, l2 = l2, l2.Next
                        rst = rst.Next
                }
        }
        return head.Next
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := add(l1, l2)
	return  fix(head)
}


