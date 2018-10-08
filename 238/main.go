package main

import(
	"fmt"
)

func productExceptSelf(nums []int) []int {
		l := len(nums)
		if l == 0{
			return nil
		}
		if l == 1 {
			return nums
		}
		lp := make([]int, l)
		rp := make([]int, l)

		lp[0] = 1
		rp[l-1] = 1

		for i := 1;i < l;i++{
			lp[i] = lp[i-1]*nums[i-1]
		}
		for i := l - 2;i >= 0;i--{
            rp[i] = rp[i+1]*nums[i+1]
        }

		rst := []int{}

		for i := 0;i < l;i++{
			rst = append(rst, lp[i] * rp[i])
		}
		return rst
}

func main(){
	in := []int{1,2,3,4}
	fmt.Println(productExceptSelf(in))
}

