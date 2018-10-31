package main

import(
		"fmt"
)


type ListNode struct {
    Val int
    Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []*ListNode
	var s []int
	len1, len2 := 0, 0
	for cur := l1; cur != nil; cur = cur.Next{
		s1 = append(s1, cur)
		len1++
	}
	for cur := l2; cur != nil; cur = cur.Next{
        s2 = append(s2, cur)
		len2++
    }
	if len1 < len2{
		len1, len2, s1, s2 = len2, len1, s2, s1
	}
	for i, carry := 0, 0;;i++{
		if i > len2 {
			for  {
				if(i > len1) {
					break
				}
				sum := s1[i].Val + carry
				v := sum % 10
				carry = sum / 10
				s = append(s, v)
			}
			if carry == 1 {
				s = append(s, carry)
				len1++
			}
			break
		}
		fmt.Println(s1[i].Val, s2[i].Val, i)
		sum := s1[i].Val + s2[i].Val + carry
		v := sum % 10 
		carry = sum / 10
		s = append(s, v)
	}
	res := &ListNode{
			Val : s[0],
	}
	now := res
	for i := 1; i <= len1; i++{
			now.Next = &ListNode{
            	Val : s[i],
    		}
			now = now.Next
	}
	return res
}

func newListNode(a []int) *ListNode{
	res := &ListNode{
		Val : a[0],
	}
	now := res
	l := len(a)
	for i:=1;i<l;i++{
			now.Next = &ListNode{
    		    Val : a[i],
    		} 
			now = now.Next
	}
	return res
}

func fn(l *ListNode){
	for l != nil {
		fmt.Println(l)
		l = l.Next
	}
}


func main(){
	a1 := []int{
		7,2,4,3,
	}
	a2 := []int{
		5,6,4,
	}
	l1, l2  := newListNode(a1), newListNode(a2)
	fn(l1)
	fn(l2)
	res := addTwoNumbers(l1, l2)
	fn(res)
}
