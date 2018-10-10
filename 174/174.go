func fn(self, v int) int{
	res := v-self
	if res < 1{
		res = 1
	}
	return res
}

func min(v1, v2 int) int{
		if v1 != -1 && (v2 == -1 || v2 != -1 && v1 < v2){
			return v1
		}
		return v2
}

func calculateMinimumHP(dungeon [][]int) int {
    stI, stJ, enI, enJ := 0, 0, len(dungeon), len(dungeon[0])
	res := make([][]int, enI)
	for i := range(res){
		res[i] = make([]int, enJ)
	}
	enI--
	enJ--
	res[enI][enJ] = fn(dungeon[enI][enJ],1)

	for i := enI; i>=0; i--{
		for j := enJ; j>=0; j--{
			vi, vj := -1, -1
			if i >= enI && j >= enJ {
				continue
			}
			if i < enI{
				vi = res[i+1][j]
			}
			if j < enJ{
				vj = res[i][j+1]
			}
			v := min(vi, vj)
			res[i][j] = fn(dungeon[i][j], v)
		}
	}
	return res[stI][stJ]
}
