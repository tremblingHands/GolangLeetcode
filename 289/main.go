package main

import(
	"fmt"
)




func gameOfLife(board [][]int)  {
		ll, lv := len(board), len(board[0])
		fn := func(i, j int) int{
			if i<0 || i>=ll || j<0 || j>=lv{
				return 0
			}else{
				return board[i][j]%2
			}

		}
		for i, v := range(board){
				for j, _ := range(v){
					sum := fn(i-1,j-1) +  fn(i,j-1) + fn(i+1,j-1) + fn(i-1,j) +
						 fn(i+1,j) + fn(i-1,j+1) + fn(i,j+1) + fn(i+1,j+1)
					if board[i][j] == 0{
						if sum == 3{
							board[i][j] = 2
						}
					}else {
						if sum < 2 || sum > 3{
							board[i][j] = 3
						}
					}
				}
		}
		for i, v := range(board){ 
                for j, _ := range(v){
					if board[i][j] == 2{
						board[i][j] = 1
					}else if board[i][j] == 3{
						board[i][j] = 0
					}
                }      
        }
}


func main(){
		board := [][]int{
			{0,1,0},
			{0,0,1},
			{1,1,1},
			{0,0,0},
		}
		gameOfLife(board)
		for _, v := range(board){
			fmt.Println(v)
		}
}
