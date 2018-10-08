package main

import(
	"fmt"
	"sort"
)


type Point struct {
    X int
    Y int
}

var o Point
var line bool = true
type Points []Point

func (p Points) Len() int{
	return len(p)
}
func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Points) Less(i, j int) bool{
	c := (p[i].X - o.X) * (p[j].Y - o.Y) - (p[i].Y - o.Y) * (p[j].X - o.X)
	if c == 0 { 
		return (p[i].X - o.X) * (p[i].X - o.X) + (p[i].Y - o.Y) * (p[i].Y - o.Y) >
				(p[j].X - o.X) * (p[j].X - o.X) + (p[j].Y - o.Y) * (p[j].Y - o.Y)
	}else {
		line = false
		return c > 0
	}
}


func outerTrees(points []Point) []Point {
	// if len(points) <= 2
	o = points[0] 
	var que Points
	for _, p := range (points) {
		if p.Y < o.Y{
			o = p
		}
		que = append(que, p)
	}
	sort.Sort(que)
	if line == true {
		return points
	}
	que = que[:len(que)-1]
	tmp := que 
	res := make([]Point, len(points))
	res[0], res[1] = o, que[0]
	que = que[1:]
	l := 2
	for{
		if len(que) <= 0 {
			break
		}
		cur := que[0]
		i, j := l-2, l-1
		if (cur.X - res[i].X) * (res[j].Y - res[i].Y) - (cur.Y - res[i].Y) * (res[j].X - res[i].X) <= 0{
			res[l] = cur
			l++
			que = que[1:]
		}else{
			l--
		}
	}
	return res[:l]
}

func main(){
	var points []Point 
	in := [][]int{{1,1},{2,2},{2,0},{2,4},{3,3},{4,2}}
	//in := [][]int{{1,5}}
	//in := [][]int{{1,2},{2,2},{4,2}}
	//in := [][]int{{3,0}, {4,0}, {5,0}, {4,2}, {4,5}}
	for _, p := range(in){
		points = append(points, Point{
			X : p[0],
			Y : p[1],
		})
	}
	fmt.Println(points)
	fmt.Println(outerTrees(points))
}



