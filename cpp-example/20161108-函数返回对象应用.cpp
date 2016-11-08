#include <iostream>
#include <stdlib.h>

using namespace std;

class book
{
public:
    void setprice(double a);
    double getprice();
    void settitle(char* a);
    char* gettitle();
private:
    double price;
    char* title;
};
void book::setprice(double a)
{
    price = a;
}
double book::getprice()
{
    return price;
}
void book::settitle(char* a)
{
    title = a;
}
char* book::gettitle()
{
    return title;
}

void display(book& b)
{
    cout<<"The price of "<<b.gettitle()<<" is $"<<b.getprice()<<endl;
}

book& init(char* t, double p)
{
    static book b;
    b.settitle(t);
    b.setprice(p);
    cout<<"&b = "<<&b<<endl;
    return b;
}

int main() 
{ 
    book Alice;
    Alice.settitle("Alice in wonderland");
    Alice.setprice(29.9);
    display(Alice);

    //方式一:Harry是指向静态对象b的引用 
    book& Harry = init("Harry Potter", 49.9);
    cout<<"&Harry="<<&Harry<<endl; 

    //方式二:Harry2是新的处对象,init返回b对象的应用然后把b对象赋值给Harry2对象
    book Harry2;
    Harry2 = init("Harry2 Potter", 69.9);
    cout<<"&Harry2="<<&Harry2<<endl; 



    display(Harry);  
     
    return 0;
}  

