
Trie树的基本性质：

根节点不包含字符，除根节点外的每一个子节点都包含一个字符。
从根节点到某一个节点，路径上经过的字符连接起来，为该节点对应的字符串。
每个节点的所有子节点包含的字符互不相同。
通常在实现的时候，会在节点结构中设置一个标志，用来标记该结点处是否构成一个单词（关键字）。

可以看出，Trie树的关键字一般都是字符串，而且Trie树把每个关键字保存在一条路径上，而不是一个结点中。另外，两个有公共前缀的关键字，在Trie树中前缀部分的路径相同，所以Trie树又叫做前缀树（Prefix Tree）。

##优缺点

Trie树的核心思想是空间换时间，利用字符串的公共前缀来减少无谓的字符串比较以达到提高查询效率的目的。
优点
插入和查询的效率很高，都为$O(m)$，其中 $m$ 是待插入/查询的字符串的长度。

关于查询，会有人说 hash 表时间复杂度是$O(1)$不是更快？但是，哈希搜索的效率通常取决于 hash 函数的好坏，若一个坏的 hash 函数导致很多的冲突，效率并不一定比Trie树高。
Trie树中不同的关键字不会产生冲突。

Trie树只有在允许一个关键字关联多个值的情况下才有类似hash碰撞发生。

Trie树不用求 hash 值，对短字符串有更快的速度。通常，求hash值也是需要遍历字符串的。

Trie树可以对关键字按字典序排序。

缺点
当 hash 函数很好时，Trie树的查找效率会低于哈希搜索。

空间消耗比较大。

##Trie树的应用

1、字符串检索
2、词频统计
3、字符串排序
4、前缀匹配
5、作为其他数据结构和算法的辅助结构
如后缀树，AC自动机等。

```
#include <iostream>
#include <string>
using namespace std;

#define ALPHABET_SIZE 26

typedef struct trie_node
{
	int count;   // 记录该节点代表的单词的个数
	trie_node *children[ALPHABET_SIZE]; // 各个子节点 
}*trie;

trie_node* create_trie_node()
{
	trie_node* pNode = new trie_node();
	pNode->count = 0;
	for(int i=0; i<ALPHABET_SIZE; ++i)
		pNode->children[i] = NULL;
	return pNode;
}

void trie_insert(trie root, char* key)
{
	trie_node* node = root;
	char* p = key;
	while(*p)
	{
		if(node->children[*p-'a'] == NULL)
		{
			node->children[*p-'a'] = create_trie_node();
		}
		node = node->children[*p-'a'];
		++p;
	}
	node->count += 1;
}

/**
 * 查询：不存在返回0，存在返回出现的次数
 */ 
int trie_search(trie root, char* key)
{
	trie_node* node = root;
	char* p = key;
	while(*p && node!=NULL)
	{
		node = node->children[*p-'a'];
		++p;
	}
	
	if(node == NULL)
		return 0;
	else
		return node->count;
}

int main()
{
	// 关键字集合
	char keys[][8] = {"the", "a", "there", "answer", "any", "by", "bye", "their"};
	trie root = create_trie_node();

	// 创建trie树
	for(int i = 0; i < 8; i++)
		trie_insert(root, keys[i]);

	// 检索字符串
	char s[][32] = {"Present in trie", "Not present in trie"};
	printf("%s --- %s\n", "the", trie_search(root, "the")>0?s[0]:s[1]);
	printf("%s --- %s\n", "these", trie_search(root, "these")>0?s[0]:s[1]);
	printf("%s --- %s\n", "their", trie_search(root, "their")>0?s[0]:s[1]);
	printf("%s --- %s\n", "thaw", trie_search(root, "thaw")>0?s[0]:s[1]);

	return 0;
}
```

