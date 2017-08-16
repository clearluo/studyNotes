/*
1017. A除以B (20)
本题要求计算A/B，其中A是不超过1000位的正整数，B是1位正整数。你需要输出商数Q和余数R，使得A = B * Q + R成立。

输入格式：

输入在1行中依次给出A和B，中间以1空格分隔。

输出格式：

在1行中依次输出Q和R，中间以1空格分隔。

输入样例：
123456789050987654321 7
输出样例：
17636684150141093474 3
*/
#include <stdio.h>
#include <string.h>
#define MAX 1000
int main(){
    char data[MAX+1];
    int b, i, item, yu = 0;
    //freopen("a.in", "r", stdin);
    memset(data, 0x00, sizeof(data));
    scanf("%s %d", data, &b);
    for (i=0; data[i]!='\0'; i++){
        item = data[i]-'0';
        yu = yu *10 + item;
        if (yu / b == 0){
            if(i!=0 || strlen(data)<2){
                printf("%d", 0);
            }
        }else{
            printf("%d", yu / b);
            yu %= b;
        }
    }
    printf(" %d", yu);
    return 0;
}
