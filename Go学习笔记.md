<h1 align="center">Go学习笔记<h1>

## 一.基础语法

### 基础类型和运算符

1. s:=123这种定义是int32还是int64还是其他类型？试着赋值给其给uint32或int64看是否需要强制转换? s:=123这种定义是int类型,需要强制转换才能赋值给uint32或uint64

2. 在 fmt.Printf 中使用下面的说明符来打印有关变量的相关信息：
    * %+v 打印包括字段在内的实例的完整信息
    * %#v 打印包括字段和限定类型名称在内的实例的完整信息
    * %T 打印某个类型的完整说明

3. 在Go语言中,&& 和 || 是具有快捷性质的运算符，当运算符左边表达式的值已经能够决定整个表达式的值的时候（&& 左边的值为 false，|| 左边的值为 true），运算符右边的表达式将不会被执行。利用这个性质，如果你有多个条件判断，应当将计算过程较为复杂的表达式放在运算符的右侧以减少不必要的运算。

4. 格式化说明
  在格式化字符串里，%d 用于格式化整数（%x 和 %X 用于格式化 16 进制表示的数字），%g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法），%0d 用于规定输出定长的整数，其中开头的数字 0 是必须的。%n.mg 用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 或者 f，例如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00。

5. 运算符于优先级
  有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

  | 优先级  |                运算符                 |
  | :--: | :--------------------------------: |
  |  7   |               ^    !               |
  |  6   | *    /    %    <<    >>    &    &^ |
  |  5   |         +    -    \|    ^          |
  |  4   |   ==    !=    <    <=    >=    >   |
  |  3   |                 <-                 |
  |  2   |                 &&                 |
  |  1   |                \|\|                |
6. 获取字符串中某个字节的地址的行为是非法的,例如:&str[i]
### 函数

#### 闭包

匿名函数同样被称之为闭包(函数时语言的术语):它们被允许调用定义在其它环境下的变量.闭包可使得某个函数扑捉到一些外部状态,例如:函数被创建时的状态.另一种表示方式为:一个闭包继承了函数所声明时的作用域.这种状态(作用域内的变量)都被共享到闭包的环境中,因此这些变量可以在闭包中被操作,直到被销毁

```go
package main

import "fmt"

func main() {
	f := colosure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
func colosure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

/*
//运行结果
0xc0420381d0
0xc0420381d0
11
0xc0420381d0
12
0xc0420381d0
13
*/

```

#### Defer

使用defer实现代码追踪:
```go
package main

import "fmt"

func main() {
	b()
}
func trace(s string) string {
	fmt.Println("Entering:", s)
	return s
}
func un(s string) {
	fmt.Println("Leaving:", s)
}
func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}
func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

/*
//运行结果
Entering: b
in b
Entering: a
in a
Leaving: a
Leaving: b
*/

```

使用defer语句来记录函数的参数与返回值

```go
package main

import (
	"io"
	"log"
)

func main() {
	func1("Go")
}
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v\n", s, n, err)
	}()
	return 7, io.EOF
}

/*
//运行结果
2017/05/05 17:01:54 func1("Go") = 7, EOF
*/

```

```go
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

/*
//运行结果
3
3
3
2
1
0
*/

```

#### Panic,Recover

beego中api接口每个携程有安装recover函数,不必自己安装,也就是说如果接口里面panic了,主进程不回挂掉,而如果接口里面自己又go func自己创建了携程,则此时携程里面如果panic会导致整个进程挂掉,避免的方式是在子携程函数里面再安装recover,如果此时panic,不会影响到主进程.

```go
package main

import "fmt"

func main() {
	A()
	B()
	C()
}
func A() {
	fmt.Println("Func A")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}
func C() {
	fmt.Println("Func C")
}

/*
//运行结果
Func A
Recover in B
Func C
*/
```

```go
package main

import "fmt"

func main() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_colosure i = ", i) }()
		fs[i] = func() { fmt.Println("colosure i = ", i) }
	}
	for _, f := range fs {
		f()
	}
}

/*
//运行结果
colosure i =  4
colosure i =  4
colosure i =  4
colosure i =  4
defer_colosure i =  4
defer i =  3
defer_colosure i =  4
defer i =  2
defer_colosure i =  4
defer i =  1
defer_colosure i =  4
defer i =  0
*/
```

#### 变长参数

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	mytest(a...) // 此时为引用传递
	fmt.Println(a)
	x, y, z := 11, 22, 33
	mytest(x, y, z) // 此时为值传递
	fmt.Println(x, y, z)
}
func mytest(a ...int) {
	for i := range a {
		a[i] = 33
	}
}

/*
//运行结果
[1 2 3]
[33 33 33]
11 22 33
*/
```
func test(a ...interface{})传递验证

```go
package main

import "fmt"

func main() {
	mytest(1, "aa", 88.6)
}
func mytest(a ...interface{}) {
	fmt.Println(a)
}

/*
//运行结果
[1 aa 88.6]
*/
```


#### 工厂函数

一个返回值为另一个函数的函数可以被称之为工厂函数，这在您需要创建一系列相似的函数的时候非常有用：书写一个工厂函数而不是针对每种情况都书写一个函数。下面的函数演示了如何动态返回追加后缀的函数：

```go
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
```

现在,我们可以生成如下函数:

```go
addBmp := MakeAddSuffix(".bmp")
addJpeg := MakeAddSuffix(".jpeg")
//然后调用它们：
addBmp("file") // returns: file.bmp
addJpeg("file") // returns: file.jpeg
```

可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，是函数式语言的特点。我们已经在第 6.7 中得知函数也是一种值，因此很显然 Go 语言具有一些函数式语言的特性。闭包在 Go 语言中非常常见，常用于 goroutine 和管道操作

### 数组与切片

#### 基础概念
* 数组长度也是数组类型的一部分，所以[5]int和[10]int是属于不同类型的。数组的编译时值初始化是按照数组顺序完成的

* 元素的数目，也称为长度或者数组大小必须是固定的并且在声明该数组时就给出（编译时需要知道数组长度以便分配内存）；数组长度最大为 2Gb

* 如果我们想让数组元素类型为任意类型的话可以使用空接口作为类型当使用值时我们必须先做一个类型判断

* 由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。如果要使用指针，那么就需要用到后面介绍的slice类型了。验证函数传递

* ```go
  package main

  import "fmt"

  func main() {
  	arr := [...]int{1, 2, 3}
  	fmt.Println("all before:", arr)
  	test_value(arr)
  	fmt.Println("value after:", arr)
  	test_point(&arr)
  	fmt.Println("point after:", arr)
  	test_slice(arr[:])
  	fmt.Println("slice after:", arr)
  }
  func test_value(arr [3]int) {
  	arr[1] = 111
  	fmt.Println("test_value:", arr)
  }
  func test_point(arr *[3]int) {
  	arr[1] = 222
  	fmt.Println("test_point:", arr)
  }
  func test_slice(arr []int) {
  	arr[1] = 333
  	fmt.Println("test_slice:", arr)
  }

  ```

* 注意:绝对不要用指针指向slice. 切片本身已经是一个引用类型,所以它本身就是一个指针.

* []interface{}这种切片是否可以存储不同数据类型? 可以,因为该切片出存储的类型是interface{}
#### new和make的区别

* 看起来二者没什么区别,都在堆上分配内存,但是它们的行为不同个,适用于不同的类型;

* new(T)为每个新的类型T分配一片内存,初始化为0并且返回类型为*T的内存地址:这种方法返回一个指向类型为T,值为0的地址的指针,它适用于值类型,如数组和结构体;它相当于&T{};

* make(T)返回一个类型为T的初始值,它只适用于3种内建的引用类型:切片,map和channel;

* 换言之,new函数分配内存,make函数初始化;

* 下面的例子说明了在映射上使用 new 和 make 的区别以及可能发生的错误：

  ```go
  package main

  type Foo map[string]string
  type Bar struct {
  	thingOne string
  	thingTwo int
  }

  func main() {
  	// OK
  	y := new(Bar)
  	(*y).thingOne = "hello"
  	(*y).thingTwo = 1
  	// NOT OK
  	z := make(Bar) // 编译错误：cannot make type Bar
  	(*z).thingOne = "hello"
  	(*z).thingTwo = 1
  	// OK
  	x := make(Foo)
  	x["x"] = "goodbye"
  	x["y"] = "world"
  	// NOT OK
  	u := new(Foo)
  	(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
  	(*u)["y"] = "world"
  }
  ```
* 试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，但是 new() 一个映射并试图使用数据填充它，将会引发运行时错误！ 因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎。

### Map

### 指针

### 包

### 结构体

### 接口

### 反射

### Goroutine和Channel

## 二.深入理解

### defer,return相关
1. defer、return、返回值之间执行顺序的坑
  Go语言中延迟函数defer充当着 try...catch 的重任，使用起来也非常简便，然而在实际应用中，很多gopher并没有真正搞明白defer、return和返回值之间的执行顺序，从而掉进坑中，今天我们就来揭开它的神秘面纱！
  先来运行下面两段代码：

  A. 匿名返回值的情况

  ```go
  package main
  import (
  "fmt"
  )
  func main() {
  fmt.Println("a return:", a()) // 打印结果为 a return: 0
  }
  func a() int {
  var i int
  defer func() {
  	i++
  	fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
  }()
  defer func() {
  	i++
  	fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
  }()
  return i
  }
  ```
  B. 有名返回值的情况
  ```go
  package main
  import (
  "fmt"
  )
  func main() {
  fmt.Println("b return:", b()) // 打印结果为 b return: 2
  }
  func b() (i int) {
  defer func() {
  	i++
  	fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
  }()
  defer func() {
  	i++
  	fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
  }()
  return i // 或者直接 return 效果相同
  }
  ```
  先来假设出结论（这是正确结论），帮助大家理解原因：

  * 多个defer的执行顺序为“后进先出”；
  * 所有函数在执行RET返回指令之前，都会先检查是否存在defer语句，若存在则先逆序调用defer语句进行收尾工作再退出返回；
  * 匿名返回值是在return执行时被声明，有名返回值则是在函数声明的同时被声明，因此在defer语句中只能访问有名返回值，而不能直接访问匿名返回值；
  * return其实应该包含前后两个步骤：第一步是给返回值赋值（若为有名返回值则直接赋值，若为匿名返回值则先声明再赋值）；第二步是调用RET返回指令并传入返回值，而RET则会检查defer是否存在，若存在就先逆序插播defer语句，最后RET携带返回值退出函数；
  * 因此，defer、return、返回值三者的执行顺序应该是：return最先给返回值赋值；接着defer开始执行一些收尾工作；最后RET指令携带返回值退出函数。

  如何解释两种结果的不同：

  * 上面两段代码的返回结果之所以不同，其实从上面的结论中已经很好理解了。
  * a()int 函数的返回值没有被提前声名，其值来自于其他变量的赋值，而defer中修改的也是其他变量（其实该defer根本无法直接访问到返回值），因此函数退出时返回值并没有被修改。
  * b()(i int) 函数的返回值被提前声名，这使得defer可以访问该返回值，因此在return赋值返回值 i 之后，defer调用返回值 i 并进行了修改，最后致使return调用RET退出函数后的返回值才会是defer修改过的值。

  C. 下面我们再来看第三个例子，验证上面的结论：

  ```go
  package main
  import (
  "fmt"
  )
  func main() {
  c:=c()
  fmt.Println("c return:", *c, c) // 打印结果为 c return: 2 0xc082008340
  }
  func c() *int {
  var i int
  defer func() {
  	i++
  	fmt.Println("c defer2:", i, &i) // 打印结果为 c defer2: 2 0xc082008340
  }()
  defer func() {
  	i++
  	fmt.Println("c defer1:", i, &i) // 打印结果为 c defer1: 1 0xc082008340
  }()
  return &i
  }
  ```
  虽然 c()\*int 的返回值没有被提前声明，但是由于 c()\*int 的返回值是指针变量，那么在return将变量 i 的地址赋给返回值后，defer再次修改了 i 在内存中的实际值，因此return调用RET退出函数时返回值虽然依旧是原来的指针地址，但是其指向的内存实际值已经被成功修改了。
  即，我们假设的结论是正确的！

2. 思考以下示例,函数f返回时,变量ret的值是什么?

  ```go
  package main
  import "fmt"
  func main() {
  fmt.Println(f())
  }
  func f() (ret int) {
  defer func() {
  	ret++
  }()
  return 1
  }
  /*
  变量ret的值为2,因为ret++是在执行return 1语句以后执行的,也就是说defer是在return以后才执行,这可用于在返回语句之后修改返回的error时使用
  */
  ```

### 切片相关

1. 通过切片引用传递给函数,在函数中修改此切片后,调用者得到的切片一定是函数中修改后的内容吗?

   未必,通过以此达到返回数据给调用者是不安全的,因为如果在此函数中操作切片内容大于原有切片的空间的话,将从新分配空间,此时调用者使用的还是原来数组控件(即便在此函数中操作切片内容不大于原有切片空间,调用者和函数中的切片也不是同一个,他们地址相同,但是长度不同)

   ```go
   package main

   import "fmt"

   func main() {
   	a := []byte{0, 1, 2}
   	fmt.Println(a)
   	mytest(a)
   	fmt.Println(a)
   }
   func mytest(a []byte) {
   	a[0] = 00
   	a[1] = 11
   	a[2] = 22
   	a = append(a, 55)
   	a[0] = 77
   	fmt.Println("cap(a) = ", cap(a))
   }

   // 运行结果
   // [0 1 2]
   // cap(a) =  8
   // [0 11 22]
   ```

### 函数相关
1. 不使用递归但使用闭包实现斐波那契数列程序

   ```go
   package main

   import "fmt"

   func main() {
   	fmt.Println(mytest(9))
   }
   func mytest(num int) int {
   	f := fibonacci()
   	for i := 0; i < num; i++ {
   		f()
   	}
   	return f()
   }
   func fibonacci() func() int {
   	back1, back2 := 1, 1
   	return func() int {
   		temp := back1
   		back1, back2 = back2, (back1 + back2)
   		return temp
   	}
   }

   // 运行结果
   // 55
   // 可使用数组存储每次计算的fibonacci的值,不必每次计算
   ```




## 三.官方标准库
