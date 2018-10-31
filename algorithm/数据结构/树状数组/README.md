//树状数组（Binary Indexed Tree（BIT），Fenwick Tree）是一个查询和修改复杂度都为 log(n) 的数据结构，主要用于查询任意两位之间的所有元素之和，但每次只能修改一个元素的值。
//让第 i 个位置管理[i-lowbit(i)+1, i]这一段区间


// 获取 x 的二进制中为 1 的最后一位对应的十进制，也即 x 能被整除的最大的2的幂 
int lowbit(int x){
	return x & (-x);
}

void modify(int x, int add){
	while(x <= MAXN){
		a[x] += add;
		x += lowbit(x);
	}
}

int get_sum(int x){
	int ret = 0;
	while(x != 0){
		ret += a[x];
		x -= lowbit(x);
	}	
	return ret;
}

void modify(int x, int y, int data){
	for(int i = x; i < MAXN; i += lowbit(x)){
		for(int j = y; j < MAXN; j += lowbit(j))
			a[i][j] += data;
	}
}

int get_sum(int x, int y){
	int res = 0;
	for(int i = x; i > 0; i -= lowbit(i)){
		for(int j = y; j > 0; j -= lowbit(j){
			res += a[i][j];
		}
	}
	return res;
}





