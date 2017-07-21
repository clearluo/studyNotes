/*
1002. 写出这个数 (20)
读入一个自然数n，计算其各位数字之和，用汉语拼音写出和的每一位数字。

输入格式：每个测试输入包含1个测试用例，即给出自然数n的值。这里保证n小于10100。

输出格式：在一行内输出n的各位数字之和的每一位，拼音数字间有1 空格，但一行中最后一个拼音数字后没有空格。

输入样例：
1234567890987654321123456789
输出样例：
yi san wu

 */
#include <stdio.h>
int main()
{
    int i;
    int c;
    int sum = 0;
    int data[100];
    char * str[]={"ling","yi","er","san","si","wu","liu","qi","ba","jiu"};
    //freopen("a.in", "r", stdin);
    while ((c = getchar()) != EOF) {
        if (c >= '0' && c <= '9'){
            sum += (c-'0');
        }
    }
    i = 0;
    if (sum == 0) {
        printf("ling");
        return 0;
    }
    while (sum > 0){
        data[i] = sum % 10;
        sum /= 10;
        i++;
    }
    for (i = i-1; i>=0; i--){
        printf("%s", str[data[i]]);
        if (i != 0) {
            printf(" ");
        }
    }
    return 0;
}