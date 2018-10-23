package main

import(
		"fmt"
)

type node struct{
	k,v int
	pre, next *node
}

type Chain struct{
	head *node
	tail *node
}

type LRUCache struct {
    capacity int
	length int
	table map[int]*node
	chain Chain
}


func Constructor(capacity int) LRUCache {
	tmp := &node{}
	return LRUCache{
		capacity,
		0,
		make(map[int]*node),
		Chain{
			head: tmp,
			tail: tmp,
		},
	}
}


func (this *LRUCache) Get(key int) int {
		if n, o := this.table[key];!o{
			fmt.Println(-1)
			return -1
		}else{
			if n != this.chain.tail{
				n.next.pre = n.pre
				n.pre.next = n.next
				this.chain.tail.next = n
				n.pre = this.chain.tail
				n.next = nil
				this.chain.tail = n
			}
			fmt.Println(n.v)
			return n.v
		}
}


func (this *LRUCache) Put(key int, value int)  {
		if n, o := this.table[key];o{
			n.v = value
			this.Get(key)
			return
		}
		in := &node{
			v:value,
			k:key,
		}
    	if this.length<this.capacity{
			in.pre = this.chain.tail
			this.chain.tail.next = in
			this.chain.tail = in
			this.table[key] = in
			this.length++
		}else{
			out := this.chain.head.next
			delete(this.table, out.k)
			this.chain.head.next = out.next
			if out.next != nil{
				out.next.pre =  this.chain.head
			}
			if out.next == nil{
				this.chain.tail = this.chain.head
			}
			in.pre = this.chain.tail
            this.chain.tail.next = in
            this.chain.tail = in
            this.table[key] = in
		}

}

var obj LRUCache 

func pp(){
	now := obj.chain.head
	for{
		now = now.next
		if now == nil{
			break
		}
		fmt.Print(now.v)
	}
	fmt.Println()
}
func main(){
		capacity := 1
		obj = Constructor(capacity)
		obj.Put(2,1)
		obj.Get(2)
		obj.Put(3,2)
		obj.Get(2)
		obj.Get(3)
}
