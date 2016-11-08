#include <iostream>
#include <stdlib.h>

using namespace std;

class Book
{
public:
    void setprice(double a);
    double getprice();
private:
    double price;
};
void Book::setprice(double a)
{
    price = a;
}
double Book::getprice()
{
    return price;
}

void display(Book& b)
{
    cout<<"The price is $"<<b.getprice()<<endl;
}

Book& init(double p)
{
    static Book b;
    b.setprice(p);
    cout<<"static &b = "<<&b<<endl;
    return b;
}

int main() 
{ 
    //方式一:book_reference是指向静态对象b的引用 
    Book& book_reference = init(49.9);
    cout<<"&book_reference="<<&book_reference<<endl; 

    //方式二:book_object是新的对象,init返回b对象的引用然后把b对象赋值给book_object对象
    Book book_object;
    book_object = init(69.9);
    cout<<"&book_object="<<&book_object<<endl; 

    display(book_reference);  
    display(book_object);


    book_reference.setprice(33.9);

    book_object.setprice(99.9);

    cout<<"=======book_reference引用对象========="<<endl;
    display(book_reference);

    cout<<"=======book_object对象========="<<endl;
    display(book_object);
     
    return 0;
}  

