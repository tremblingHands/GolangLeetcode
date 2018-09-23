package main

import(
	"fmt"
	"strconv"
	"math"
)

func op(s string) (int, error) {
	l := len(s)
	t := make([]byte, l)
	copy(t, s)
	for i ,j := 0, l-1; i < j;{
		t[j] = t[i]
		i++
		j--
	}
	return strconv.Atoi(string(t))
}

func change(s string ,v int) string{
	l := len(s)
	id := 0
	if l%2 == 0{
		id = l/2
	}else{
		id = l/2+1
	}
	n, _ := strconv.Atoi(s[0:id])
	n += v
	rst := strconv.Itoa(n)
	if rst == "0"{
		rst = ""
	}
	if len(rst) < id{
        	return strconv.Itoa(int(math.Pow10(l-1)-1))
        }
	return rst + s[id:]
}

func Abs(n int) int {
	if n < 0{
		n = -n
	}
	return n
}

func main(){
	in := ""
	fmt.Scanf("%s", &in)
	l := len(in)
	if l == 1{
		one, _ := strconv.Atoi(in)
		fmt.Println(strconv.Itoa(one-1))
		return
	}
	mid, _ := op(in)
	pre, _ := op(change(in, -1))
	post, _ := op(change(in, 1))

	fmt.Println("now", mid, pre,post)

	n, _ := strconv.Atoi(in)
	min := Abs(pre-n)
	minV := pre
	if min > Abs(mid-n) && Abs(mid-n) != 0 {
		min, minV  = Abs(mid-n),  mid
	}
	if min > Abs(post-n) {
		min, minV = Abs(post-n),  post
	}
	fmt.Println(strconv.Itoa(minV))

}

