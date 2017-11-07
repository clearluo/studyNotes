/*
1004. 成绩排名 (20)
读入n名学生的姓名、学号、成绩，分别输出成绩最高和成绩最低学生的姓名和学号。

输入格式：每个测试输入包含1个测试用例，格式为

  第1行：正整数n
  第2行：第1个学生的姓名 学号 成绩
  第3行：第2个学生的姓名 学号 成绩
  ... ... ...
  第n+1行：第n个学生的姓名 学号 成绩
其中姓名和学号均为不超过10个字符的字符串，成绩为0到100之间的一个整数，这里保证在一组测试用例中没有两个学生的成绩是相同的。
输出格式：对每个测试用例输出2行，第1行是成绩最高学生的姓名和学号，第2行是成绩最低学生的姓名和学号，字符串间有1空格。

输入样例：
3
Joe Math990112 89
Mike CS991301 100
Mary EE990830 95
输出样例：
Mike CS991301
Joe Math990112
 */
#include <stdio.h>
#include <string.h>

int main()
{
    int n;
    char hightName[11] = {0};
    char hightNo[11] = {0};
    int hightScore = -1;
    char lowName[11] = {0};
    char lowNo[11] = {0};
    int lowScore = 101;
    char tempName[11] = {0};
    char tempNo[11] = {0};
    int tempScore = 0;

    //freopen("a.in", "r", stdin);
    scanf("%d",&n);
    while(n--){
        memset(tempName,0x00,sizeof(tempName));
        memset(tempNo,0x00,sizeof(tempNo));
        scanf("%s %s %d",tempName,tempNo,&tempScore);
        if (tempScore > hightScore){
            strcpy(hightName,tempName);
            strcpy(hightNo,tempNo);
            hightScore = tempScore;
        }
        if (tempScore < lowScore){
            strcpy(lowName,tempName);
            strcpy(lowNo,tempNo);
            lowScore = tempScore;
        }
    }
    printf("%s %s\n",hightName,hightNo);
    printf("%s %s\n",lowName,lowNo);
    return 0;
}