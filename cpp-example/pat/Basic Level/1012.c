/*
1012. 数字分类 (20)
给定一系列正整数，请按要求对数字进行分类，并输出以下5个数字：

A1 = 能被5整除的数字中所有偶数的和；
A2 = 将被5除后余1的数字按给出顺序进行交错求和，即计算n1-n2+n3-n4...；
A3 = 被5除后余2的数字的个数；
A4 = 被5除后余3的数字的平均数，精确到小数点后1位；
A5 = 被5除后余4的数字中最大数字。
输入格式：

每个输入包含1个测试用例。每个测试用例先给出一个不超过1000的正整数N，随后给出N个不超过1000的待分类的正整数。数字间以空格分隔。

输出格式：

对给定的N个正整数，按题目要求计算A1~A5并在一行中顺序输出。数字间以空格分隔，但行末不得有多余空格。

若其中某一类数字不存在，则在相应位置输出“N”。

输入样例1：
13 1 2 3 4 5 6 7 8 9 10 20 16 18
输出样例1：
30 11 2 9.7 9
输入样例2：
8 1 2 4 5 6 7 9 16
输出样例2：
N 11 2 N 9

*/
#include <stdio.h>
#include <math.h>

int main(){
    int A1 = 0;
    int A2 = 0, countA2 = 0;
    int A3 = 0;
    float A4 = 0.0, sum = 0.0, countA4 = 0.0;
    int A5 = 0;
    int n, temp, remainder;
    //freopen("a.in","r",stdin);
    scanf("%d", &n);
    while (n--) {
        scanf("%d", &temp);
        remainder = temp % 5;
        switch(remainder){
        case 0:
            if (temp % 2 == 0){
                A1 += temp;
            }
            break;
        case 1:
            A2 = A2 + (temp) * pow(-1,countA2);
            countA2++;
            break;
        case 2:
            A3++;
            break;
        case 3:
            sum += temp;
            countA4++;
            break;
        case 4:
            if (temp > A5){
                A5 = temp;
            }
            break;
        }
    }

    A1 == 0 ? printf("N ") : printf("%d ", A1);

    countA2 == 0 ? printf("N ") : printf("%d ", A2);

    A3 == 0 ? printf("N ") : printf("%d ", A3);

    countA4 == 0 ? printf("N ") : printf("%.1f ", sum / countA4);

    A5 == 0 ? printf("N") : printf("%d", A5);

    return 0;
}
