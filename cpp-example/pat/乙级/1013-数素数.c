/*
1013. 数素数 (20)
令Pi表示第i个素数。现任给两个正整数M <= N <= 104，请输出PM到PN的所有素数。

输入格式：

输入在一行中给出M和N，其间以空格分隔。

输出格式：

输出从PM到PN的所有素数，每10个数字占1行，其间以空格分隔，但行末不得有多余空格。

输入样例：
5 27
输出样例：
11 13 17 19 23 29 31 37 41 43
47 53 59 61 67 71 73 79 83 89
97 101 103
*/
#include <stdio.h>
#include <math.h>

int isPrimer(x){
    int i;
    int k = sqrt(x);
    for (i=2; i<=k; i++){
        if (x % i == 0){
            return 0;
        }
    }
    return 1;
}
int main(){
    int m, n, x = 1;
    int count = 0, realCount = 0;
    //freopen("a.in","r",stdin);
    scanf("%d %d", &m, &n);
    while(x++){
       if (!isPrimer(x)){
        continue;
       }
       count++;
       if (count>=m && count<=n){
            if (realCount % 10 != 0){
                printf(" ");
            }
            printf("%d", x);
            if (realCount % 10 == 9){
                printf("\n");
            }
            realCount++;
       }
       if (count > n){
        break;
       }
    }
    return 0;
}
