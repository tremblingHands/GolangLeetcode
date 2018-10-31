package main

import(
	"fmt"
)

func lowbit(x int) int{
	return x & (-x)
}


func update(in []int, pos int){
		l := len(in)
		for{
			if pos >= l{
				break
			}
			in[pos]++
			pos += lowbit(pos)
		}
}

func getsum(in []int, pos int) int{
		sum := 0
		for{
			if pos <= 0{
					break
			}
			sum += in[pos]
			pos -= lowbit(pos)
		}
		fmt.Println(pos, sum)
		return sum
}


func main(){
	//in := []int{9,1,0,5,4}
	pos := []int{5,2,1,4,3}
	bt := make([]int, 6)
	sum := 0
	for i:=1;i<=5;i++{
		update(bt, pos[i-1])
		sum += (pos[i-1]-getsum(bt, pos[i-1]))
	}
	fmt.Println(sum)

}

