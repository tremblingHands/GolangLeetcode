
并查集一般用于对动态连通性的判断，主要应用于判断两个元素是否在同一个集合，两个点是否连通，变量名等同性以及间接好友的判断，同时并查集经常作为其他模板的一部分实现某些功能

并查集常用于的题型为判断某两个元素是否属于同一个集合，判断图是否连通或是否有环，或配合其他算法如最小生成树Kruskal，与DP共同使用等。

1,初始化一个并查集 initUnionFind
初始化并查集,一般是将每个元素作为一个单独的集合,对于下面的示例就是对应的元素父节点就是自己
2,合并两个不相交集合 union
两个元素,分别找到(find)他们的根节点,然后将其中一个元素的根节点的父亲指向另外的一个元素的根节点
3,查找某个元素的根节点 find
查找一个元素的根节点,parent--->parent--->parent.....
那么问题来了,查找元素根节点途径的元素过多,那么就有一个优化的问题-------路径压缩,直接将该元素的父亲指向根节点或者祖先
4,判断两个元素是否属于同一个集合 isConnected
判断两个元素是否属于同一个集合,就是分别找到他们的根节点,然后判断两个根节点是否相等.


```
private int count;
    private int[] parents;
    
    //初始化并查集
    public void initUnionFind(int m, int n, char[][] grid){
        parents = new int[m*n];
        for(int i=0; i<m; i++){
            for(int j=0; j<n; j++){
                if(grid[i][j] == '1'){
                    count++;
                }
                parents[i*n+j] = i*n+j;
            }
        }
    }
    
    public int find(int idx){
        while(idx != parents[idx]){
            //在查找的过程中压缩路径,减少查找的次数
            parents[idx] = parents[parents[idx]];
            idx = parents[idx];
        }
        return idx;
    }

	public int find(int idx){
		if parents[idx] != idx{
			parents[idx] = find(parents[idx]);
		}
		return parents[idx];
	} 

    public void union(int p, int q){
        int pRoot = find(p);
        int qRoot = find(q);
        //两个元素的根不同,则合并
        if(pRoot != qRoot){
            parents[qRoot] = pRoot;
            count--;
        }
    }
    
    public boolean isConnected(int p, int q){
        int pRoot = find(p);
        int qRoot = find(q);
        
        //两点不连通
        if(pRoot != qRoot){
            return false;
        }
        return false;
    }
```
