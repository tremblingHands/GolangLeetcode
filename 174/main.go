package main

import(
	"fmt"
)

var st, enI, enJ, minPower int


func dfs(power, min , i, j int, dungeon [][]int){
	if i > enI || j > enJ{
		return
	}
	if min > minPower{
		return
	}
	if i == enI && j == enJ{
		power = power + dungeon[i][j]
		if power * -1 > min{
        	min = power * -1
    	}
		if minPower > min{
			minPower = min
		}
		return
	}
	power = power + dungeon[i][j]
	if power * -1 > min{
		min = power * -1
	}
	fmt.Println(i,j,  power, min)
	dfs(power, min, i+1, j, dungeon)
	dfs(power, min, i, j+1, dungeon)

}

func calculateMinimumHP(dungeon [][]int) int {
	enI = len(dungeon)-1
	enJ = len(dungeon[0])-1
	minPower = 1<<32
	min := -1<<32
	dfs(0, min, 0, 0, dungeon)
	if minPower <0 {
        minPower = 0
    }
	return minPower+1
}

func main(){
	/**
	dungeon := [][]int{
		{-2,-3,3},
		{-5,-10,1},
		{10,30,-5},
	}
	**/
	dungeon := [][]int{
        {0,0,0},
		{1,1,-1},
    }
	fmt.Println(calculateMinimumHP(dungeon))
}


