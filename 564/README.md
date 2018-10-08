Given an integer n, find the closest integer (not including itself), which is a palindrome.

The 'closest' is defined as absolute difference minimized between two integers.

Example 1:
Input: "123"
Output: "121"
Note:
The input n is a positive integer represented by string, whose length will not exceed 18.
If there is a tie, return the smaller one as answer.

#solution

Palindrome 是指前半部分和后半部分相同
因此，距离原始字符串最近的Palindrome出现在下面三种情况：

1. 原始字符串的前半部分不变，根据前半部分构造Palindrome的后半部分
2. 原始字符串的前半部分数值+1，再根据前半部分构造Palindrome的后半部分
3. 原始字符串的前半部分数值-1，再根据前半部分构造Palindrome的后半部分

需要考虑
1. 距离原始字符串最近的Palindrome不包括自身
2. 前半部分以“9”结尾的+1；前半部分以“0”结尾的-1

证明：
前半部分不变、+1、-1都可以产生唯一Palindrome，且距离最近
