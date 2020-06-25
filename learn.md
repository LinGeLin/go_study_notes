# go 有点

- 生产力

复合

# 与其他主要编程语言的差异

1、go 中main函数不支持任何返回值，通过os.Exit返回状态

2、main函数不支持传入参数，在程序中直接通过os.Args获取命令行参数。

# go语言的常量和变量

## 编写测试程序

1、源码文件以 _test 结尾：xxx_test.go

2、测试方法名以 Test 开头： func TestXXX(t *testing.T) {...}

> 大写的开头代表包外可以访问

## 变量赋值与其他主要语言的差异

1、赋值可以自动类型推断

2、在一个赋值语句中可以对多个变量进行同时赋值，可以大大简化代码。

```go
a, b = b, a
```

## 常量

- 快速设置连续值

```go
const (
    Monday = itoa + 1  // 1
    TuesDay            // 2
    Wednesday          // ....
    Thursday
    Friday
    Saturday
    Sunday
)

const (
    Open = 1 << itoa // 001
    Close            // 010
    Pending          // 100
)
```

# go 数据类型

go 基本数据类型 |
---|
bool | 
string |
int int8 int16 int32 int64 |
uint uint uint16 uint32 uint64 uintptr |
byte // alias(别名) for uint8 |
rune // alias for int32， represents(表现) a Unicode code point |
float32 float64 |
complex64 complex128 |

> 不允许隐式类型转换，原类型与别名也不支持

```go
type MyInt int64
var int64 b = 3
var MyInt c
c = b //不被允许
```

错误代码，这是一段测试go语言类型转换的代码
```go
package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64 = 3
	b = (int64)a
	var c MyInt = 4
	// c = b
	t.Log(a, b, c)
}
```
报错代码
```go
b = (int64)a
```
改正
```go
b = int64(a)
```
go语言版本：`1.14.4`


## 指针类型

1、不支持指针运算

2、string  是值类型，其默认的初始化值为空字符串，长度为0，而不是nil

# 运算符

1、go语言中没有前置的++，--，只支持后置++，--

2、用 == 比较数组

- 相同维数且含有相同个数的数组才可以比较

- 每个元素都相同的才相等

3、按位清零运算符 &^

```go
1 &^ 0 ---1
1 &^ 1 ---0
0 &^ 1 ---0
0 &^ 0 ---0

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry1(t *testing.T) {
    a := 7 //0111
    a = a&^ Readable // 清除读功能
    a = a&^ Writable // 清除写功能
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
```

# 条件和循环

go语言仅支持 for 循环

```go
n := 0

for n < 5 {
    n++
    fmt.Println(n)
}

// 无限循环

n := 0
for {
    ...
}
```

if 条件与其他编程语言的差异
```go
if condition {
    ...
} else {

}

if condition - 1 {
    ...
} else if condition - 2 {
    ...
} else {
    ...
}
```

1、condition 表达式结果必须为布尔值

2、支持变量赋值
```go
if var declaration; condition {
    ...
}
```

switch 条件 

1、条件表达式不限制为常量或者整数；

2、单个 case 中，可以出现多个结果选项，使用逗号分隔；

3、与C语言等规则相反，Go语言不需要用 break 来明确退出一个case；

4、可以不设定 switch 之后的条件表达式，在此种情况下，整个switch结构与多个 if...else...的逻辑作用等同。

case 语句更像 if else 语句。

在某些写法上相对其他语言有一些简化。

```go
switch os := runtime.GOOS; os {
    case "darwin":
        ...
    case "linux":
        ...
    default:
        ...
}


switch {
    case 0 <= Num && Num <=3:
        ...
    case 4 <= Num && Num <=6:
        ...
    case 7 <= Num && Num <=9:
        ...
}
```

# go 语言中常用的集合

## 数组和切片

- 数组的声明

```go
var a [3]int // 声明并初始化为默认值零
a[0] = 1

b := [3]int{1,2,3} // 声明同时初始化
c := [2][2]int{{1,2}, {3,4}} // 多为数组初始化

//遍历

for _, e := range arr { // _表示占位
    t.Log(e)
}

for idx, e := range arr {
    t.Log(idx, e)
}
```
- 数组的截取（切片）

a[开始索引（包含）: 结束索引（不包含）]

```go
a := [...]int{1, 2, 3, 4, 5}
a[1:2] // 2
a[1, len(a)] // 2, 3, 4, 5
a[1:] // 2, 3, 4, 5
a[:3] // 1, 2, 3
```

- 切片内部结构

结构 | 类型  | 含义
--- |---    | ---
ptr | *Elem | 指向一片连续的存储空间，指向一个数组 
len | int   | 元素的个数，可以访问的元素的个数
cap | int   | 内部数组的容量，大于等于len

```go
s2 := make([]int, 2, 4)
// []int type, 2 len, 4 cap
// 其中len个元素会被初始化为默认值，未初始化元素不可访问
```

切片共享存储结构，当两个切片同时指向一块内存时，通过分别通过append进行增加，会同时修改所指内存的内容，当其中的一个append操作导致capacity增加时，两个切片将不再指向同一块内存。

- 切片如何实现可变长？

```go
func TestSliceGrowing(t *testing.T) {
    s := []int{}
	for i :=0; i<10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
```
capacity 成倍增长，当存储空间去扩展的时候，会创建一个新的连续存储空间，并把原有的数值都拷贝过去。slice的自增长很方便，但是性能上需要再考虑。

months | index  
--- |---
**  | 0
Jan | 1
Feb | 2
Mar | 3
Apr | 4
May | 5
Jun | 6
Jul | 7
Aug | 8
Sep | 9
Oct | 10
Nov | 11
Dec | 12

```go
Q2 = months[4:7]
// len = 3, cap = 9

Summer = months[6:9]
// len = 3, cap = 7
```

## 数组与切片的不同

1、容量是否可伸缩，数组不可以进行伸缩，切片可以进行伸缩

2、是否可以进行比较，相同长度，相同维数的数组可以进行比较，切片只能与`nil`进行比较

3、声明的时候

```go
// slice
slice = []int{1, 2, 3, 4}

// array
array = [4]int{1, 2, 3, 4}
```

# Map

## map声明
```go
m := map[string]int{"one" : 1, "two" : 2, "three" : 3}

m1 := map[string]int{}

m1["one"] = 1

m2 := make(map[string]int, 10 /* Initial Capacity*/)
// 为什么不初始化len？
// map没有办法做初始化
// map和切片都是自增长度的一种数据结构，初始化容量避免内存扩容带来的性能消耗
```

## 与其他编程语言的差异

再访问的key不存在时，仍会返回零值，不能通过返回 `nil` 来判断元素是否存在。

```go
m := map[int]int{}
if v, ok := m[3]; ok {
    fmt.Println("Exist")
} else {
    fmt.Println("Not Exist")
}
```

## map遍历

并不保证次序

```go
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1:1, 2:4, 3:9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
```

# Map 与工厂模式

- Map 的 value 可以时一个方法

- 与 Go 的 Dock type 接口方法一起，可以方便的实现单一方法对象的工厂模式

```go
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op}
	t.Log(m[1](2), m[2](2), m[3](2))
}
```

## 实现Set

可以保证添加元素的唯一性，可以判断唯一元素的个数。

go 的内置集合中没有 Set 实现，可以 map[type]bool

1、元素的唯一性

2、基本操作
    
    1、添加元素

    2、判断元素是否存在

    3、删除元素

    4、元素个数

```go
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true // 添加元素
	n := 1
	if mySet[n] { // 判断元素是否存在
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

	mySet[3] = true
	t.Log(len(mySet)) // 元素个数

	delete(mySet, 1) // 删除元素
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
```

# 字符串

- 与其他主要编程语言的差异

1、string 是数据类型，不是引用或指针类型

2、string 是只读的 byte slice， len 函数返回包含的 byte 数

3、string 的 byte 数组可以存放任何数据，可见字符，二进制字符

```go
package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认值""
	
	s = "hello"
	t.Log(len(s))
	// s[1] = '3'
	// string 是不可变的 byte slice

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "\xE4\xB8\xA5\xFF"
	t.Log(s)
}
```

- Unicode 与 UTF8

1、Unicode 是一种字符集 (code point)

2、UTF8 是 Unicode 的存储实现（转换为字节序列的规则）

字符 | "中"
--- | ---
Unicode | 0x4E2D
UTF-8 | 0xE4B8AD
string / []byte | [0xE4, 0xB8, 0xAD]

- 常用的字符串函数

1、strings 包

2、strconv 包

```go
package string_test

import (
	"testing"
	"strings"
	"strconv"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
```

# 函数 （一等公民）

- 与其他主要编程语言的差异

1、可以又多个返回值

2、所有参数都是值传递：slice、map、channel 会有传引用的错觉，切片是包含指针的数据结构。

3、函数可以作为变量的值，eg函数可以作为map的值

4、函数可以作为参数和返回值

```go
package func_test

import (
	"math/rand"
	"testing"
	"time"
	"fmt"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}


func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
```

- 函数可变参数及defer（延迟运行）

可变长参数，并不需要指定参数的个数，但是其类型一致。将参数转化为数组，通过遍历数组获取参数

```go
func sum(ops ...int) int {
    s :=0
    for _, op := range ops {
        s += op
    }
    return s
}
```

defer 函数，可用来释放资源或者释放锁

```go
func TestDefer(t *testing.T) {
    defer func() {
        t.Log("Clean resources")
    }()
    t.Log("Started")
    panic("Fatal error") // defer 仍会执行，程序异常中断
}
```

# go面向对象

是不是面向对象？ Yes and no，是也不是，go 不支持继承

- 封装数据和行为

结构体

```go
type Employee struct {
    Id string
    Name string
    Age int
}
```
实例创建及初始化

```go
e := Employee{"0", "Bob", 20}

e1 := Employee{Name : "Mike", Age : 30}

e2 := new(Employee) //这里返回的引用/指针，相当于 e := &Employee{}
e2.Id = "2" // 通过实例的指针访问成员不需要使用->
e2.Age = 22
e2.Name = "Rose"
// %T输出变量的类型
// 访问结构体的数据，无论是不是指针都通过 "." 符号
```
行为（方法）定义与其他主要编程语言的差异

```go
type Employee struct {
    Id string
    Name string
    Age int
}

// 第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制，有较大空间开销
func(e Employee) String() string {
    fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
    return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String() string {
    fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
    return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
    e := Employee{"0", "Bob", 20}
    // e := &Employee{"0", "Bob", 20}，也可以正常执行
    // 为什么指针也可以调用？
    // 通过一个类型的指针的实例调用成员或者方法不需要通过箭头->
    // 所以 e.String() 可以调用到正常的方法
	t.Log(e.String())
}
```

# Go 语言相关接口

// 循环依赖

```java
// Programmer.java

public interface Programmer {
	String WriteCodes();
}

// GoProgrammer.java

public class GoProgrammer implements Programmer {
	@Override
	public String WriteCodes() {
		return "fmt.Println(\"Hello World\")";
	}
}

// Task.java
public class Task {
	public static void main(String[] args) {
		Programmer prog = new GoProgrammer();
		String codes = prog.WriteCodes();
		System.out.println(codes);
	}
}
```

- go 语言的interface （Duck Type式接口实现）

接口定义 

```go
type Programmer interface {
	WriteHelloWorld() Code
}
```

接口实现

```go
package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

func (g * GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new (GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
```

1、接口为非入侵性，实现不依赖与接口定义，

2、所有接口的定义可以包含在接口使用者包内

- 接口变量（函数类型变量）

```go
var prog Coder = &GoProgrammer{}

type GoProgrammer struct {

}

&GoProgrammer
```
prog | 初始化之后包含两部分
---- |---
类型 | type GoProgrammer struct
数据 | &GoProgrammer

- 自定义类型(既可以是函数，也可以是基本类型)

```go
type IntConvertionFn func(n int) int

type MyPoint int
```

# 扩展与复用（面向对象）

go 不支持继承

```go
package extension

import (
	"testing"
	"fmt"
)

type Pet struct {

}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
		// 内嵌的类型不支持重载和lsp
	Pet // 匿名嵌套类型
}

// func (d *Dog) Speak() {
// 	fmt.Print("wang")
// }

// func (d *Dog) SpeakTo(host string) {
// 	d.Speak()
// 	fmt.Println(" ", host)
// }

func TestDog(t *testing.T) {
	// var dog Pet := new(Dog) 编译错误，不支持隐式类型转换
	// 
	// var dog *Dong = new(Dog) // lsp
	// var p = (*Pet)(dog)
	// p.SpeakTo("Chao")
	// 编译错误
	dog := new(Dog)
	dog.SpeakTo("Chao")
}

// 输出 ... Chao
```

# 多态

```go
type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}
// 
type GoProgrammer struct {

}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

//
type JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

func writeFistProgram(p Programmer) { // Programmer是interface，所以传进的参数只能是指针
	fmt.Printf("%T, %v\n", p, p.WriteHelloWorld())// %T 输出类型
}

func TestPloymoriphsim(t *testing.T) {
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFistProgram(goProg)
	writeFistProgram(javaProg)
}
```
> 输出
> 1. %v    只输出所有的值
>
> 2. %+v 先输出字段类型，再输出该字段的值
>
> 3. %#v 先输出结构体名字值，再输出结构体（字段类型+字段的值）

```go
package main
import "fmt"
type student struct {
	id   int32
	name string
}


func main() {
	a := &student{id: 1, name: "xiaoming"}
 
	fmt.Printf("a=%v	\n", a)
	fmt.Printf("a=%+v	\n", a)
	fmt.Printf("a=%#v	\n", a)
}

a=&{1, xiaoming}
a=&{id:1, name:xiaoming}
a=&main.student{id:1, name:xiaoming}
```

# 空接口与断言

1、空接口可以表示任何类型（c中的void *）

2、通过断言来将空接口转换为指定类型（并非通过强制类型转化）

```go
v, ok := p.(int) // ok=true时转换成功，自己定义的时候要符合这种习惯
```

```go
package empty_interface

import (
	"testing"
	"fmt"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if i, ok := p.(string); ok {
	// 	fmt.Println("String", i)
	// 	return 
	// }
	// fmt.Println("Unknow Type")

	switch v:=p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("Hello World")
}
```

- go 接口最佳实践

1、倾向于使用小的接口定义，很多接口只包含一个方法(倡导更小的接口，负担更轻 signal method interface)
```go
type Reader interface {
	Read(p []byte)(n int, err error)
}

type Writer interface {
	Write(p []byte)(n int, err error)
}
```

2、较大的接口定义，可以由多个小接口定义组合而成
```go
type ReadWriter interface {
	Reader
	Writer
}
```

3、只依赖于必要功能的最小接口
```go
func StoreData(reader Reader) error {

}
```

# go的错误机制

与其他主要编程语言的差异

1、没有异常机制

2、error 类型实现error接口

3、可以通过 errors.New 来快速创建错误实例

(通过函数的第二返回值判断运行正确或错误)

```go
type error interface {
	Error() string
}

errors.New("n must be in the range[0, 100]")
```

- 最佳实践

1、定义不同的错误变量，以便判断错误类型
```go
package error_test

import (
	"testing"
	"errors"
)

var LessThanTwoError = errors.New("n should not be less than 2")
var LargerThanHunderedError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHunderedError
	}
	fibList := []int{1, 1}
	for i:=2; i<n; i++ {
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1001); err != nil {
		if err == LessThanTwoError {
			t.Error(err)
		} else if (err == LargerThanHunderedError) {
			t.Error(err)
		}
	} else {
		t.Log(v)
	}
}
```

2、及早失败，避免嵌套

```go
func GetFibonacci1(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Print(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

// 推荐下面这种方法，及早退出，避免嵌套
func GetFibonacci2(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}

	if list,err = GetFibonacci(i); err != nil {
		fmt.Println("Error, err")
		return
	}
	fmt.Println(list)
}
```
# 错误处理之 panic 和 recover

- panic

1、panic 用于不可回复的错误

2、panic 退出前会执行 defer（延迟函数） 指定的内容

- panic vs os.Exit

1、os.Exit 退出时不会调用 defer（延迟函数） 指定的函数

2、os.Exit 退出时不输出当前调用栈信息，而 panic 会将函数调用栈打印出来

- recover

Java
```java
try {
	...
} catch(Throwable t) {

}
```
C++
```C++
try {
	...
} catch(...) {

}
```
go
```go
defer func() {
	// recover会返回一个错误，这个错误就是 panic 传进去的错误
	// 然后针对错误进行恢复
	if err := recover(); err != nil {
		// 恢复错误
	}
}()
```
```go
package panic_recover

import (
	"errors"
	// "os"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
	fmt.Println("End")
}

// print 
Start
recovered from Something wrong
```

常见的"错误恢复"，具有危险性
```go
defer func() {
	if err := recover(); err != nil {
		log.Error("recovered panic", err)
	}
}()
```
- 当心 recover 成为恶魔

1、形成僵尸服务进程，导致health check 失效（可能是系统资源消耗完了，但是强制恢复过来，程序仍无法正常进行服务，大多health check只是检查进程在还是不在）

2、“let it Crash！”往往时我们恢复不确定性错误的最好方法（进程Crash之后，守护进程对服务进程进行重启）

# 构建可复用的模块（包）

1、基本复用模块单元

> 以首字母大写来表明可被包外代码访问，反之不可被访问


2、代码的 package 可以和所在的目录不一致（java目录名和包名统一成一致的）

3、同一目录里 Go 代码的 package 要保持一致

`GOPATH`
```go
export GOPATH="/home/lin/Desktop/golearn"
```
`serise`
```go
package series

// 开头是小写，则包外无法访问
func square(n int) int {
	return n*n
}

func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-1] + ret[i-2])
	}
	return ret
}
```
`client`
```go
package client

import (
	"testing"
	"ch15/series"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	t.Log(series.square(1)) //出错，小写开头的无法访问
}
```

- init 方法

1、在 main 被执行前，所有依赖的 package 的 init 方法都会被执行

2、不同包的 init 函数按照包导入的依赖关系决定执行顺序

3、每个包可以有多个 init 函数

4、包的每个源文件也可以有多个 init 函数，这点比较特殊，相同的名称

```go
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}
```
- 使用远程 package

1、通过 go get 来获取源程依赖

> go get -u 强制从网络更新远程依赖

2、注意代码在 GitHub 上的组置形式，以适应 go get

> 直接以代码路径开始，不要有 src

`ConcurrentMap for GO`

https://github.com/easierway/concurrent_map.git

```go
go get -u github.com/easierway/concurrent_map
// 没有 https:// 和 .git

package remote_package

import (
	"testing"

	// 起别名
	cm "github.com/easierway/concurrent_map"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

// result
=== RUN   TestConcurrentMap
    TestConcurrentMap: remote_package_test.go:12: 10 true
--- PASS: TestConcurrentMap (0.00s)
PASS
ok      ch15/remote_package     0.001s
```

# 依赖管理

- Go 未解决的依赖问题

1、同一环境下，不同项目使用同一包的不同版本（没有办法使用不同版本，package 在get下来之后都被放在 GOPATH 路径下，每个 project 都会按照 go 指定的依赖管理去查找 package，先找 GOPATH ，然后找 GOROOT, 就没有办法把指定版本的 package 放到指定目录下面）

2、无法管理对包的特定版本的依赖

- vender 路径

随着 Go 1.5 release 版本的发布，vender 目录被添加到被添加到除了 GOPATH 和 GOROOT 之外的依赖目录查找的解决方案，在 Go 1.6 之前，你需要手动设置环境变量

**查找依赖包路径的解决方案如下：**

1、当前包下的 `vender` 目录

2、向上级目录查找，直到找到 `src` 下的 `vender` 目录

3、在 `GOPATH` 下面查找依赖包

4、在 `GOROOT` 目录下查找

多了一个自己 project 的 vender 目录，去放我们依赖的 package

- 常见的依赖管理工具

```go
godep 

glide

dep

// glide

sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
sudo apt-get install glide

glide init	// 初始化
glide install // 生成vender目录
```

# 协程机制

- Thread vs Groutine

1、创建时默认的 stack 的大小

JKD5 以后 Java Thread stack 默认为1M

Groutine 的 Stack 初始化大小为 2K (创建速度更快)

2、和 KSE （Kernel Space Entity）的对应关系（创建的线程/协程与内核线程对象的比例）

Java Thread 是 1:1

Groutine 是 M:N （多对多）

一比一线程进行切换的时候会牵涉到内核线程的切换，多对多则不会

user space thread scheduler |.
--- |---
thread | groutine
thread | thread1  thread2  thread3
kernel space entity scheduler |
kernel entity | kernel entity

MPG |. |
--- | --- 
M | System Thread 系统线程
P | Processor 处理器，Go语言实现的协程处理器
G | Goroutine 协程

Processor 在不同的系统线程中，每个 Processor 都挂着一个准备运行的协程队列，有一个协程处于正在运行状态， Processor 依次运行队列中的协程。

如果一个协程的运行时间特别长，把整个 Processor 都占着了，那么在队列中的协程是不是都会延迟很久呢？

答：在 go 起来的时候，会有一个守护线程，会去做一个计数，统计每个 Processor 运行完成的协程的数量，一段时间后发现发现某个 Processor 完成的协程数量没有发生变化的时候，就会往协程的任务栈中插入一个特别的标记，当协程运行的时候遇到非内联函数的时候就会读到这个特别的标记，就会把自己中断下来，把自己插到等待协程队列的队尾，切换下一个协程运行。

- 另一个提高并发的能力的机制

当某一个协程被系统中断了，比如等待系统 IO 时，为了提高整体的并发，Processor 会把自己移到另一个可以使用的系统线程中，继续执行它所挂队列里的其他的协程。当被中断协程被唤醒了，它会把自己加到某一个 Processor 的协程等待队列中，或者加到全局等待队列中。

当协程被中断的时候，在寄存器中的状态也会被保存在协程对象中，当协程再次获得运行机会的时候，这些又会被重新写入寄存器，继续运行。

```go
package groutine_test

import (
	"testing"
	"fmt"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i< 10; i++ {
		// go 的方法调用都是值传递，传递i的时候都复制了一份
		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)

		go func () {
			// i 是共享的，共享就存在竞争条件
			fmt.Println(i)
		}()
		// 输出十行10
	}
	time.Sleep(time.Millisecond * 50)
}
```

# 共享内存并发机制

```java
Lock lock = ...;
lock.lock();
try {
	// process (thread-safe)
} catch(Exception ex) {

} finally {
	lock.unlock();
}
```

go 语言

```go
package sync

Mutex

RWLock //读写锁
```

```go
package share_mem

import (
	"testing"
	"time"
	"sync"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		// 没有并发的保护，最终输出counter将小于5000
		go func() {
			counter ++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		// 有并发保护，最终输出 counter 等于 5000
		go func() {
			defer func () {
				mut.Unlock()
			}()
			mut.Lock()
			counter ++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}
```

**WaitGroup（类似java中的join）**

只有当我wait的所有的东西都执行完之后才能继续往下执行
```go
var wg sync.WaitGroup
for i := 0; i < 5000; i++ {
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		...
	}()
}
wg.Wait()
```
```go
func TestCounterWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mt sync.Mutex
	counter := 0;
	for i:=0; i<5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mt.Unlock()
			}()
			mt.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
```

# CSP （Communicating sequential processes通信顺序进程）并发控制

依赖一个通道完成两个通信实体之间的协调

**CSP vs Actor**

1、和 Actor 的直接通讯不同，CSP 模式则是通过 Channel 进行通讯的，更松耦合一些

2、Go 中 channel 是有容量限制并且独立与处理 Groutine，而如 Erlang，Actor模式中的 messagebox 是无限的，接收进程也总是被动地处理消息，go有限制。

**channel**

1、通讯的两方必须同时在 channel 上，任何一方不在的时候另一方都会被阻塞在哪里等待。如，往 channel 内放消息一方会因为接受消息的一方未就绪而阻塞。
```go
// 会阻塞直到取到值
retCh = make(chan string)
```

2、buffer channel。channel 有一定的容量，放消息的人在 channel 没满的情况下可以一直放消息，接收消息的人在 channel 不为空的情况下可以一直从 channel 中读取消息。
```go
// 不会阻塞
retCh = make(chan string, 1)
```

一段java代码
```java
private static FutureTask<String> service() {
	FutureTask<String> task = new FutureTask<String>(()->"Dosomething");
	new Thread(task).start();
	return task;
}

FutureTask<String> ret = service();
System.out.println("Do something else");
System.out.println(ret.get()); // 差不多准备好了，获取结果，节省的是准备的时间
```
go代码
```go
package csp

import (
	"testing"
	"time"
	"fmt"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)
	// retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)// 需要结果时从channel中取出结果
}
// 

working on something else
returned result.
Task is done.
Done
```
# 多路选择和超时

**select**
```go
select {
	case ret := <- retCh1:
		t.Logf("result %s", ret)
	case ret := <- retCh2:
		t.Logf("result %s", ret)
	default:
		t.Error("No one returned")
}
```
**超时控制**
```go
select {
	case ret := <- retCh:
		t.Logf("result %s", ret)
	case <- time.After(time.Second * 1):
		t.Errot("time out")
}
```

```go
package select_test

import (
	"testing"
	"time"
	"fmt"
)

func service() string {
	time.Sleep(time.Millisecond * 120)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string)
	// retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <- AsyncService():
		t.Log(ret)
	case <- time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
//
=== RUN   TestSelect
    TestSelect: select_test.go:49: time out
--- FAIL: TestSelect (0.10s)
FAIL
exit status 1
```
# channel的关闭与广播

**channel 的关闭**

1、向关闭的 channel 发送数据，会导致 panic

2、`v , ok <-ch;` ok 为 bool 值，true 表示正常接收，false 表示通道关闭

3、所有的 channel 接收者都会在 channel 关闭时，立刻从阻塞等待中返回且上述 ok 值为 false。这个广播机制常被利用，进行向多个订阅者同时发送信号，如：退出信号。

4、channel 已经关闭了，再从 channel 中取值，则会立即返回一个通道类型的 “零值”

```go
package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// 在不确定要放多少数字的时候，如何通知接收者已经放完了？
		// eg：放token（如-1），但是reveiver会很多，而producer不知道其数量
		// 
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 发送数据完毕之后将channel关闭掉
		close(ch)
		// 关闭之后再往 channel 上发数据，会 panic
		// ch <- 11
		wg.Done()
	}()
}


// 消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// 如果通道已经关闭，再从通道中取值
		// 会立即返回通道类型的“零值”
		// 下例多输出一个0
		// for i := 0; i < 11; i++ {
		// 	data := <-ch
		// 	fmt.Println(data)
		// }

		for {
			// ok 返回 false 的时候表示 channel 已经关闭了
			if data, ok := <- ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	wg.Wait()
}
```

# 任务的取消

```go
package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <- cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	// 向 channel 发送一个消息，取消一个任务
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	// 关闭 channel 取消所有任务
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if (isCancelled(cancelCh)) {
					break;
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	//cancel_1(cancelChan)
	// 输出一行

	cancel_2(cancelChan)
	// 输出五行
	time.Sleep(time.Second * 1)
}
```

# Context(上下文) 与任务取消

管理任务的取消

|||Main||||
---|---|---|---|---|---
||Handle(Req1)||Handle(Req2)|
Search(A)|Search(B)|Search(C)|Search(A)|Search(B)|Search(C)

取消中间层次节点的时候，其关联节点也需要取消。

**Context**

1、根 Context：通过 context.Background() 创建

2、子 Context：context.WithCancel(parentContext) 创建

	ctx, cancel := context.WithCancel(context.Background())

	cancel 方法，用来取消当前 context 执行的 task

	ctx 用来传入子任务，当前 context 被取消时， ctx 关联的子任务都会被取消

3、当前 Context 被取消时，基于他的子 context 都会被取消

4、接收取消通知 <-ctx.Done()

```go
package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	 select {
	 case <- ctx.Done():
		return true
	 default:
		return false
	 }
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if (isCancelled(ctx)) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	// 调用 cancel 方法，取消当前 context 关联的所有节点
	cancel()
	time.Sleep(time.Second * 1)
}
```
# 典型并发任务

## 仅执行一次

**单例模式（懒汉式，线程安全）**

double check
```java
public class Singleton {
	private static Singleton INSTANCE=null;
	private Singleton(){}
	public static Singleton getInstance() {
		if (INSTANCE == null) {
			synchronized(Singleton.class) {
				if (INSTANCE == null) {
					INSTANCE = new Singleton()
				}
			}
		}
		return INSTANCE;
	}
}
```
go
```go
var once sync.Once
var obj *SingletonObj

func GetSingletonObj() *SingletonObj {
	// 每次调用的时候可以保证只调用一次
	once.Do(func() {
		fmt.Println("Create Singleton obj.")
		obj = & SingletonObj{}
	})
}
```
测试代码
```go
package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
	data string
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	// once 可以确保代码只执行一次
	// 所以无需去判断 singleInstance 是否为空
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			// 输出的地址是一样的，表明只创建了一次对象
			wg.Done()
		}()
	}
	// 等待所有的协程执行完毕
	wg.Wait()
}
```