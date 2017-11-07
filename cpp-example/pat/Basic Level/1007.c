/*
1007. 素数对猜想 (20)
让我们定义 dn 为：dn = pn+1 - pn，其中 pi 是第i个素数。显然有 d1=1 且对于n>1有 dn 是偶数。“素数对猜想”认为“存在无穷多对相邻且差为2的素数”。

现给定任意正整数N (< 105)，请计算不超过N的满足猜想的素数对的个数。

输入格式：每个测试输入包含1个测试用例，给出正整数N。

输出格式：每个测试用例的输出占一行，不超过N的满足猜想的素数对的个数。

输入样例：
20
输出样例：
4
 */
#include <stdio.h>
#include <math.h>
int isPrimer(int x);
int main()
{
    int n, before = 0, after = 0, count = 0;
    scanf("%d", &n);
    while (n >= 2) {
        if (isPrimer(n)) {
            after = before;
            before = n;
            if ((after - before) == 2) {
                count++;
            }
        }
        n--;
    }
    printf("%d\n",count);
    return 0;
}

int isPrimer(int x){
    int i, k;
    k = sqrt(x);
    for (i=2; i<=k; i++){
        if (x % i == 0){
            return 0;
        }
    }
    return 1;
}





