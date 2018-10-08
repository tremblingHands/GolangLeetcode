import(
	"sort"
	"container/heap"
)

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
		l := len(Profits)
		q := make(list, 0, l)
		for i := 0; i<l ; i++ {
			q = append(q, project{
					Profits[i],
					Capital[i],
			})
		}
		sort.Sort(q)
		qp := make(pro, 0, l) 
		for i := 0; i<k ; i++{
			id := 0
			for {
				if id<len(q) && q[id].capital <= W {
					heap.Push(&qp, q[id])
					id++
				}else{
					break
				}
			}
			q = q[id:]
			if len(qp) > 0 {
                W = W + heap.Pop(&qp).(project).profit 
            }else{
                break
            }
		}
		return W

}

type project struct{
	profit int
	capital int
}

type list []project
func (p list) Len() int{
	return len(p)
}
func (p list) Swap(i, j int){
	p[i], p[j] = p[j], p[i]
}
func (p list) Less(i, j int) bool{
	return p[i].capital < p[j].capital 
}

type pro []project
func (p pro) Len() int{
    return len(p)
}           
func (p pro) Swap(i, j int){
    p[i], p[j] = p[j], p[i]
}
func (p pro) Less(i, j int) bool{
    return p[i].profit > p[j].profit 
}       
func (p *pro) Push(x interface{}){
	*p = append(*p, x.(project))
}
func (p *pro) Pop() interface{}{
	l := len(*p)
	tmp := (*p)[l-1]
	*p = (*p)[:l-1]
	return tmp 
}
