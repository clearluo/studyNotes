/*
1016. 部分A+B (15)
正整数A的“DA（为1位整数）部分”定义为由A中所有DA组成的新整数PA。例如：给定A = 3862767，DA = 6，则A的“6部分”PA是66，因为A中有2个6。

现给定A、DA、B、DB，请编写程序计算PA + PB。

输入格式：

输入在一行中依次给出A、DA、B、DB，中间以空格分隔，其中0 < A, B < 1010。

输出格式：

在一行中输出PA + PB的值。

输入样例1：
3862767 6 13530293 3
输出样例1：
399
输入样例2：
3862767 1 13530293 8
输出样例2：
0
*/
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#define MAX 12
int main(){
    char A[MAX];
    char B[MAX];
    char resultA[MAX];
    char resultB[MAX];
    int indexA = 0, indexB = 0;
    char a, b;
    int i;
    //freopen("a.in", "r", stdin);
    memset(A, 0x00, sizeof(A));
    memset(B, 0x00, sizeof(B));
    memset(resultA, 0x00, sizeof(resultA));
    memset(resultB, 0x00, sizeof(resultB));
    scanf("%s %c %s %c", A, &a, B, &b);
    for (i=0; A[i]!='\0'; i++){
        if (A[i] == a){
            resultA[indexA++] = a;
        }
    }
    for (i=0; B[i]!='\0'; i++){
        if (B[i] == b){
            resultB[indexB++] = b;
        }
    }
    printf("%d", atoi(resultA) + atoi(resultB));
    return 0;
}
