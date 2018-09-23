func nearestPalindromic(n string) string {
    l := len(n)
	if l == 1{
        n, _ := strconv.Atoi(n)
		return strconv.Itoa(n-1)
	}
	mid, _ := op(n)
	pre, _ := op(change(n, -1))
	post, _ := op(change(n, 1))
	println(mid, pre,post)
	num, _ := strconv.Atoi(n)
	min := Abs(pre-num)
	minV := pre

	if min > Abs(mid-num) && Abs(mid-num) != 0 {
		min, minV  = Abs(mid-num),  mid
	}
	if min > Abs(post-num) {
		min, minV = Abs(post-num),  post
	}
	return strconv.Itoa(minV)
}

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



