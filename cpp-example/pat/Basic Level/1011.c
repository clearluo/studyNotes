/*
1011. A+B和C (15)
给定区间[-231, 231]内的3个整数A、B和C，请判断A+B是否大于C。

输入格式：

输入第1行给出正整数T(<=10)，是测试用例的个数。随后给出T组测试用例，每组占一行，顺序给出A、B和C。整数间以空格分隔。

输出格式：

对每组测试用例，在一行中输出“Case #X: true”如果A+B>C，否则输出“Case #X: false”，其中X是测试用例的编号（从1开始）。

输入样例：
4
1 2 3
2 3 4
2147483647 0 2147483646
0 -2147483648 -2147483647
输出样例：
Case #1: false
Case #2: true
Case #3: true
Case #4: false
 */
#include <stdio.h>
int main()
{
    int n, i=0;
    long int a, b, c;
    //freopen("a.in", "r", stdin);
    scanf("%d", &n);
    while(i++<n){
       scanf("%ld %ld %ld", &a, &b, &c);
       if (a + b > c){
        printf("Case #%d: true\n", i);
       } else{
        printf("Case #%d: false\n", i);
       }
    }
    return 0;
}
