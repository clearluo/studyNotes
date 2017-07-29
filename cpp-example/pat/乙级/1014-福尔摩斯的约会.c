/*
1014. 福尔摩斯的约会 (20)
大侦探福尔摩斯接到一张奇怪的字条：“我们约会吧！ 3485djDkxh4hhGE 2984akDfkkkkggEdsb s&hgsfdk d&Hyscvnm”。
大侦探很快就明白了，字条上奇怪的乱码实际上就是约会的时间“星期四 14:04”，
因为前面两字符串中第1对相同的大写英文字母（大小写有区分）是第4个字母'D'，代表星期四；
第2对相同的字符是'E'，那是第5个英文字母，代表一天里的第14个钟头（于是一天的0点到23点由数字0到9、以及大写字母A到N表示）；
后面两字符串第1对相同的英文字母's'出现在第4个位置（从0开始计数）上，代表第4分钟。
现给定两对字符串，请帮助福尔摩斯解码得到约会的时间。

输入格式：

输入在4行中分别给出4个非空、不包含空格、且长度不超过60的字符串。

输出格式：

在一行中输出约会的时间，格式为“DAY HH:MM”，其中“DAY”是某星期的3字符缩写，即MON表示星期一，TUE表示星期二，
WED表示星期三，THU表示星期四，FRI表示星期五，SAT表示星期六，SUN表示星期日。题目输入保证每个测试存在唯一解。

输入样例：
3485djDkxh4hhGE
2984akDfkkkkggEdsb
s&hgsfdk
d&Hyscvnm
输出样例：
THU 14:04

题目解析：
这题题目很坑，没有表达清楚意思；
1. 第1对除了相同的大写英文字母外，还必须符合周一到周日的字母也就是A到G；
2. 第2对相同字符是指：相同大写字母且必须是A到N，或者相同数字；
3. 第3对必须是相同英文，相同大写或相同小写；
*/
#include <stdio.h>
#include <string.h>
#define MAX 61
int main(){
    int i, index = 0;
    int first = 1, second = 1;
    int minuteIndex = 0;
    char data1[MAX];
    char data2[MAX];
    char data3[MAX];
    char data4[MAX];
    char result[3];
    char WEEK[][4] = {"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"};
    //freopen("a.in", "r", stdin);
    memset(data1, 0x00, sizeof(data1));
    memset(data2, 0x00, sizeof(data2));
    memset(data3, 0x00, sizeof(data1));
    memset(data4, 0x00, sizeof(data2));
    memset(result, 0x00, sizeof(result));
    gets(data1);
    gets(data2);
    gets(data3);
    gets(data4);
    i = 0;
    while(first || second){
        if (first){
           if(index == 0 && (data1[i] == data2[i] && data1[i]>='A' && data1[i]<='G')){
                result[index++] = data1[i];
            }else if (index == 1 && data1[i] == data2[i] && ((data1[i]>='A' && data1[i]<='N') || (data1[i]>='0' && data1[i] <= '9'))){
                result[index++] = data1[i];
            }
            if (2 == index){
                first = 0;
            }

        }
        if (second){
            minuteIndex++;
            if (data3[i] == data4[i] && ((data3[i]>='A' && data3[i]<='Z') || (data3[i]>='a' &&data3[i]<='z'))){
                result[2] = data3[i];
                second = 0;
            }
        }
        i++;
    }

    printf("%s ", WEEK[result[0]-'A']); // 是否会出现大于7的字母，是否需要除以7取余数
    if (result[1]>='0' &&result[1]<='9'){
        printf("%02d:", result[1]-'0');
    }else{
        printf("%02d:", result[1]-'A'+10);
    }

    printf("%02d", minuteIndex-1);

    return 0;
}
