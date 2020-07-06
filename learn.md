# go 

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
## 仅需任意任务完成

```go
package concurrency

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	// ch := make(chan string)
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			// 放消息，如果不是 buffer channel
			// return 之后 没有从 channel 中取消息的 receiver
			// 导致协程阻塞
			ch <- ret
		}(i)
	}
	// 当第一个人往 channel 中放消息后，接收消息的 receiver 就会从阻塞中被唤醒
	// 一旦 channel 中有消息，则会直接 return 出去
	return <- ch
}

func TestFirstRsponse(t *testing.T) {
	// 输出当前系统中的协程数
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
//
channel 
    TestFirstRsponse: fisrt_response_test.go:34: Before: 2
    TestFirstRsponse: fisrt_response_test.go:35: The result is from 4
	TestFirstRsponse: fisrt_response_test.go:37: After: 11
buffer channel
    TestFirstRsponse: fisrt_response_test.go:34: Before: 2
    TestFirstRsponse: fisrt_response_test.go:35: The result is from 4
    TestFirstRsponse: fisrt_response_test.go:37: After: 2
```

## 所有任务完成

可以使用 `sync.WaitGroup` 等待所有 `Task` 返回，也可以利用`CSP`模型。

```go
package util_all_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string  {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <- ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
```
## 对象池

将创建代价比较高的对象池化，避免重复创建，比如数据库连接、网络连接。

**使用 buffered channel 实现对象池**

归还对象 ---> 往 channel 中放数据

获取对象 ---> 从 channel 中取数据

```go
package obj_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {

}

type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓冲可重用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	// 创建 channel
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	// channel 中添加对象
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <- p.bufChan:
		return ret, nil
	case <- time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default: // 放不进去时走该分支
		return errors.New("overflow")
	}
}
```

## sync.Pool 对象缓存

**sync.Pool对象获取**

1、尝试从私有对象获取

2、私有对象不存在，尝试从当前 Processor 的共享池获取

3、如果当前 Processor 共享池也是空的，那么就尝试去其他 Processor 的共享池获取

4、如果所有子池都是空的，最后就用用户指定的 New 函数产生一个新的对象返回。

私有对象协程安全，共享池协程不安全。

**sync.Pool对象的放回**

1、如果私有对象不存在则保存为私有对象

2、如果私有对象存在，放如当前 Processor 子池的共享池中

**使用sync.Pool**

```go
// 创建池对象
pool := &sync.Pool {
	New: func() interface{} {
		return 0
	},
}

arry := pool.Get().(int)
...
pool.Put(10)
```

**sync.Pool对象的生命周期**

1、GC 会清除 sync.pool 缓存的对象

2、对象的缓存有效期为下一次 GC 之前

对象的生命周期不可知。

```go
package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool {
		New: func() interface{} {
			fmt.Println("Create a new objcet.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() // GC会清除 sync.pool 中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInMutiGroutines(t *testing.T) {
	pool := &sync.Pool {
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

**sync.Pool 总结**

1、适用于通过复用，降低复杂对象的创建和 GC 代价

2、协程安全，会有锁的开销

3、生命周期受 GC 影响，不适合于做连接池等，需自己管理生命周期的资源的池化

# 测试

## 内置单元测试框架

1、Fail, Error: 该测试失败，该测试继续，其他测试继续执行

2、FailNow, Fatal: 该测试失败，该测试终止，其他测试继续执行

```go
func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}
```

3、代码覆盖率
```go
go test -v -cover
```

4、断言
https://github.com/stretchr/testify

## Benchmark

代码片段性能测评或者第三方库性能测评。函数名 `Benchmark` 开头

执行测试

```go
// =. 或者匹配名字
go test -bench=.

// 查看内存
go test -bench. -benchmem
```

```go
func BenchmarkConcatStringByAdd(b *testing.B) {
	// 与性能测试无关的代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 测试代码
	}
	b.StopTimer()
	// 与性能测试无关的代码
}
```

```go
package benchmark_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""

	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
// 
函数名  运行次数  每一次运行所耗时间（纳秒/次）
BenchmarkConcatStringByAdd-2             7819069               139 ns/op
BenchmarkConcatStringByBytesBuffer-2    15584764                72.5 ns/op
```

## BDD （Behavior Driven Development）行为驱动开发

https://github.com/smartystreets/goconvey

安装

go get -u github.com/smartystreets/goconvery

启动 WEB UI

$GOPATH/bin/goconvey

```go
package testing
import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
```

# 反射与 Unsafe

## 反射编程

**reflect.TypeOf  vs  reflect.ValueOf**

1、`reflect.TypeOf` 返回类型 (reflect.Type)

2、`reflect.ValueOf` 返回值 (reflect.Value)

3、可以从 `reflect.Value` 获得类型

4、通过 `kind` 来判断类型

**利用反射编写灵活的代码**

1、按名字访问结构的成员

```go
reflect.ValueOf(*e).FieldByName("Name")
```

2、按名字访问结构的方法

```go
// 通过名字获取方法，Call 传入参数
reflect.ValueOf(e).MethodByName("UpdatedAge").Call([]reflect.Value{reflect.ValueOf(1)})
```

**标记 Struct Tag**

key value 结构

```go
type BasicInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
```

```go
if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
	t.Error("Failed to get 'Name' field.")
} else {
	t.Log("Tag:format", nameField.Tag.Get("format"))
}
// 
Tag:format normal
```

```go
package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

// interface{} 空接口可以传入任意类型
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)
}

type Customer struct {
	CookieID string
	Name string
	Age int
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1 : "one", 2 : "two", 3 : "three"}
	b := map[int]string{1 : "one", 2 : "two", 3 : "three"}

	t.Log("a==b?", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1==s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1==s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}

	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

type Employee struct {
	EmployeeID string
	Name string `format:"normal"`
	Age int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	// 按名字获取成员

	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}
```

## 万能程序

**DeepEqual**

切片及 map 不能进行比较，只能和 nil 进行直接比较，可以使用 reflect 中的 `DeepEqual` 进行比较

`"message": "invalid operation: a == b (map can only be compared to nil)`

**关于反射**

1、提高了程序的灵活性

2、降低了程序的可读性

3、降低了程序的性能

```go
package flexible_reflect

import (
	"errors"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1 : "one", 2 : "two", 3 : "three"}
	b := map[int]string{1 : "one", 2 : "two", 3 : "three"}

	// t.Log(a == b)
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))
}

type Employee struct {
	EmployeeID string
	Name string `format:"normal"`
	Age int
}

type Customer struct {
	CookieID string
	Name string
	Age int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {
	// func (v Value) Elem Value
	// Elem return s the value that interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr
	// It return the zero Value if v is nil
	
	// 判断是不是指针
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct type.")
	}

	// 判断是不是结构
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct{
		return errors.New("the first param should be a pointer to the struct type.")
	}

	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		// 类型相同
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			// Elem() 获取指针指向的具体的结构
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name" : "Mike", "Age" : 30}
	e := Employee{}

	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)

	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
```

## 不安全编程

主要是 c 交互的时候

**不安全行为的危险性**

```go
i := 10
// go 不支持强制类型转换，通过 unsafe.Pointer(&i)获取到指针之后可以转换为任何类型的指针
f := *(*float64)(unsafe.Pointer(&i))
```
```go
package unsafe_programming

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type Customer struct {
	Name string
	Age int
}

func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i))
	t.Log(f)
}

type MyInt int

// 合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}

// 原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}

	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFn()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 10; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()                
		}()
	}
	wg.Wait()
}
```

# 常见的架构模式的实现

## 实现 pipe-filter framework(管式过滤) 

1、非常适合数据处理及数据分析系统

2、Filter 封装数据处理的功能

3、松耦合：Filter 只跟数据（格式）耦合

4、Pipe 用于连接 Filter 传递数据或者在异步处理过程中缓冲数据流，进程内同步调用时，pipe 演变为数据在方法调用间传递

## Micro Kernel （微内核）

公共的处理流程和通用的逻辑抽象出内核，其他一些扩展的功能用作插件

**特点**

1、易于扩展

2、错误隔离

3、保持架构一致性

**要点**

1、内核包含公共流程或通用逻辑

2、将可变或可扩展部分规划为扩展点

3、抽象扩展点行为，定义接口

4、利用插件进行扩展 

# 常见任务

## 内置 JSON 解析

利用反射实现，通过 FieldTag 来标识对应的 json 值

由反射可直接得到结构域，调用结构域中的 Tag 即可获取到 tag 进行处理。`reflect.StructFied.Tag`

```go
type BasicInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type JonInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo JobInfo `json:"job_info"`
}
```

```go
package jsontest

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","c"]
	}
}	`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)

	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}
```

## easyjson

更块的 json 解析

**安装**
```go
go get -u github.com/mailru/easyjson/
```

**使用**
```go
easyjson -all struct.go
// 生成 struct_easyjson.go 文件，为对应结构添加四个方法

// 将 struct 序列化为 json
MarshJSON
MarshEasyJson

// 将 json 序列化为 struct
UnmarshalJSON
UnmarshalEasyJSON
```

**注释**
```go
//easyjson:skip 生成的时候跳过

//easyjson:json 如果没有用到参数 all 或生成

eg
//easyjson:json
type A struct {}
```

# HTTP 服务

`net/http` 标准库

**Default Router**

```go
func (sh serverHandler) ServerHTTP(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServerMux // 使用缺省的 Router
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}
	handler.ServeHTTTP(rw, req)
}
```

**路由规则**

1、URL 分为两种，末尾是 `/` 表示一个子树，后面可以跟其他子路径；末尾不是 `/` 表示一个叶子，是固定路径。

2、采用最长匹配原子，如果有多个匹配，一定采用匹配路径最长的那个进行处理

3、如果没有找到任何匹配项，会返回 404 错误

**名词解释**

1、request ：用户请求的信息，用来解析用户的请求信息，包括 post、get、cookie、url 等

2、response ：服务器返回给客户端的信息

3、conn ： 用户的每次请求链接

4、handler ： 处理请求和生成返回信息的处理逻辑

- `func StatusText(code int) string`

返回状态码所代表的信息

```go
fmt.Println(http.StatusText(200))
```

- `func CanonicalHeaderKey(s string) string`

返回头域（表示为 Header 类型）的键 s 的规范化格式，让单词首字母和 '-' 后的第一个字母大写，其余字母小写

```go
fmt.Println(http.CanonicalHeaderKey("uid-test")) // Uid-Test
fmt.Println(http.CanonicalHeaderKey("accept-encoding")) // Accept-Encoding
```
Canonical // 准确的，权威的

- `func DetectContentType(data []byte) string`

用于确定数据的 `Content-Type`， 函数总返回一个合法的 MIME 类型，如果它不能确定数据的类型，将返回 `application/octet-stream` 最多检查数据的前 512 字节

> http header 的 Content-Type 一般有三种
> application/x-www-form-urlencode：数据被编码为名称/值对，这是标准的编码格式
> multipart/form-data：数据被编码为一条消息，页上的每个控件对应消息中的一个部分
> text/plain：数据以纯文本（text/json/xml/html）进行编码，其中不含任何控件或格式字符。 postman 里标的时 RAW

```go
package main 
import(
    "fmt"
    "net/http"
)

func main() {
    cont1 := http.DetectContentType([]byte{}) //text/plain; charset=utf-8
    cont2 := http.DetectContentType([]byte{1, 2, 3}) //application/octet-stream
    cont3 := http.DetectContentType([]byte(`<HtMl><bOdY>blah blah blah</body></html>`)) //text/html; charset=utf-8
    cont4 := http.DetectContentType([]byte("\n<?xml!")) //text/xml; charset=utf-8
    cont5 := http.DetectContentType([]byte(`GIF87a`)) //image/gif
    cont6 := http.DetectContentType([]byte("MThd\x00\x00\x00\x06\x00\x01")) //audio/midi
    fmt.Println(cont1)
    fmt.Println(cont2)
    fmt.Println(cont3)
    fmt.Println(cont4)
    fmt.Println(cont5)
    fmt.Println(cont6)
}
```

- `func ParseTime(text string) (t time.Time, err error)`

用三种格式 TimeFormat，time.RFC850 和 time.ANSIC 尝试解析一个时间头的值 (Data:header)

```go
package main 
import(
    "fmt"
    "net/http"
    "time"
)


var parseTimeTests = []struct {
    h   http.Header
    err bool
}{
    {http.Header{"Date": {""}}, true},
    {http.Header{"Date": {"invalid"}}, true},
    {http.Header{"Date": {"1994-11-06T08:49:37Z00:00"}}, true},
    {http.Header{"Date": {"Sun, 06 Nov 1994 08:49:37 GMT"}}, false},
    {http.Header{"Date": {"Sunday, 06-Nov-94 08:49:37 GMT"}}, false},
    {http.Header{"Date": {"Sun Nov  6 08:49:37 1994"}}, false},
}

func main() {
    expect := time.Date(1994, 11, 6, 8, 49, 37, 0, time.UTC)
    fmt.Println(expect) //1994-11-06 08:49:37 +0000 UTC
    for i, test := range parseTimeTests {
        d, err := http.ParseTime(test.h.Get("Date"))
        fmt.Println(d)
        if err != nil {
            fmt.Println(i, err)
            if !test.err { //test.err为false才进这里
                fmt.Errorf("#%d:\n got err: %v", i, err)
            }
            continue //有错的进入这后继续下一个循环，不往下执行
        }
        if test.err { //test.err为true，所以该例子中这里不会进入
            fmt.Errorf("#%d:\n  should err", i)
            continue
        }
        if !expect.Equal(d) { //说明后三个例子的结果和expect是相同的，所以没有报错
            fmt.Errorf("#%d:\n got: %v\nwant: %v", i, d, expect)
        }
    }
}

//
userdeMBP:go-learning user$ go run test.go
1994-11-06 08:49:37 +0000 UTC
0001-01-01 00:00:00 +0000 UTC //默认返回的空值
0 parsing time "" as "Mon Jan _2 15:04:05 2006": cannot parse "" as "Mon"
0001-01-01 00:00:00 +0000 UTC
1 parsing time "invalid" as "Mon Jan _2 15:04:05 2006": cannot parse "invalid" as "Mon"
0001-01-01 00:00:00 +0000 UTC
2 parsing time "1994-11-06T08:49:37Z00:00" as "Mon Jan _2 15:04:05 2006": cannot parse "1994-11-06T08:49:37Z00:00" as "Mon"
1994-11-06 08:49:37 +0000 UTC
1994-11-06 08:49:37 +0000 GMT
1994-11-06 08:49:37 +0000 UTC
```

- `func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time`

- `func ParseHTTPVersion(ver string) (major, minor int, ok bool)`

解析 HTTP 版本字符串，如 `HTTP/1.0` 返回（1，0，true）

**header**

服务端和客户端的数据都有头部

`type Header`

```go
type Header map[string][]string

http.Header{"Date":{"1994-11-06T08:49:37Z00:00"}}
```

- `func (h Header) Get(key string) string`

返回键对应的第一个值，如果键不存在返回""

- `func (h Header) Set(key, value string)`

添加键值对到h，如键已存在则会用只有新值一个元素的切片取代旧值切片

- `func (h Header) Add(key, value string)`

添加键值对到h，如键存在，新值添加到旧值切片后面

- `func (h Header) Del(key string)`

删除键值对

- `func (h Header) Writer(w io.Writer) err`

Write 以有线格式将头域写入 w

- `func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error`

WriteSubset 以有线格式将头域写入 w，当 exclude 不为 nil 时，如果 h 的键值对的键在 exclude 中存在，且其对应值为真，该键值对就不会被写入 w

*围绕 `io.Reader/Writer` 几个常用的实现：*
>`net.Conn`，`os.Stdio`，`os.File`：网络、标准输入输出、文件的流读取
>
>`strings.Reader`：把字符串抽象成 Reader
>`bytes.Reader`：把 `[]byte` 抽象成 Reader
>`bytes.Buffer`：把 `[]byte` 抽象成 Reader 和 Writer
>`bufio.Reader/Writer`：抽象成带缓冲的流读取

**Cookie**

session 和 cookie 的区别？

session 是存储在服务器的文件，cookie 内容保存在客户端，存在被呵护篡改的情况，session 保存在服务端，可以防止被用户篡改。

**type Cookie**

Cookie 代表一个出现在 HTTP 回复的头域中 `Set-Cookie` 头里的值里，或者 HTTP 请求头域中 Cookie 头的值里
```go
type Cookie struct {
	Name		string
	Value		string
	Path		string
	Domain		string
	Expires		time.Time
	RawExpires	string
	// MaxAge = 0 表示未设置 Max-Age 属性
	// MaxAge < 0 表示立刻删除该 cookie，等价于 "Max-Age: 0"
	// MaxAge > 0 表示存在 Max-Age 属性，单位是秒
	MaxAge		int
	Secure		bool
	HttpOnly	bool
	Raw			string
	Unparsed	[]string	// 未解析的 "属性-值"对的原始文本
}
```

- `func (c *Cookie) String() string`

返回该 cookie 的序列化结果。如果只设置了 Name 和 Value 字段，序列化结果可用于 HTTP 请求的 Cookie 头或者 HTTP 回复的 Set-Cookie 头，如果设置了其他字段，则只能用于回复。

```go
Cookie := &http.Cookie{Name: "cookie-10", Value: "expiring-1601", Expires: time.Date(1601, 1, 1, 1, 1, 1, 1, time.UTC)}
Raw := "cookie-101=expiring-1601; Expires=Mon, 01 Jan 1601 01:01:01 GMT"

Cookie == Raw
```

- `func SetCookie(w ResponseWriter, cookie *Cookie)`

在 w 的头域中添加 Set-Cookie 头

常用来给 request 和 response 设置 cookie，然后使用 request 的 Cookies()、Cookie(name string) 函数和 response 的 Cookie() 函数来获取设置的 cookie 信息

**type ResponseWriter**

ResponseWrite 接口被 HTTP 处理器用于构造 HTTP 回复
```go
type ResponseWriter interface {
    // Header返回一个Header类型值，该值会被WriteHeader方法发送。
    // 在调用WriteHeader或Write方法后再改变该对象是没有意义的。
    Header() Header
    // WriteHeader该方法发送HTTP回复的头域和状态码。
    // 如果没有被显式调用，第一次调用Write时会触发隐式调用WriteHeader(http.StatusOK)
    // WriterHeader的显式调用主要用于发送错误码。
    WriteHeader(int)
    // Write向连接中写入作为HTTP的一部分回复的数据。
    // 如果被调用时还未调用WriteHeader，本方法会先调用WriteHeader(http.StatusOK)
    // 如果Header中没有"Content-Type"键，
    // 本方法会使用包函数DetectContentType检查数据的前512字节，将返回值作为该键的值。
    Write([]byte) (int, error)
}
```

**type CookieJar**

CookieJar 管理 cookie 的存储和在 HTTP 请求中的使用。CookieJar 的实现必须能安全的被多个 go 程同时使用。
```go
type CookieJar interface {
	// SetCookies 管理从u的回复中收到的cookie
	// 根据其策略和实现，它可以选择是否存储cookie
	SetCookies(u *url.URL, cookies []*Cookie)
	// Cookies 返回发送请求到u时应使用的cookie
	// 本方法有责任遵守RFC 6265规定的标准cookie限制
	Cookies(u *url.URL) []*Cookie
}
```

**type Request**

```go
type Request struct {
    // Method指定HTTP方法（GET、POST、PUT等）。对客户端，""代表GET。
    Method string
    // URL在服务端表示被请求的URI，在客户端表示要访问的URL。
    //
    // 在服务端，URL字段是解析请求行的URI（保存在RequestURI字段）得到的，
    // 对大多数请求来说，除了Path和RawQuery之外的字段都是空字符串。
    // （参见RFC 2616, Section 5.1.2）
    //
    // 在客户端，URL的Host字段指定了要连接的服务器，
    // 而Request的Host字段（可选地）指定要发送的HTTP请求的Host头的值。
    URL *url.URL
    // 接收到的请求的协议版本。本包生产的Request总是使用HTTP/1.1
    Proto      string // "HTTP/1.0"
    ProtoMajor int    // 1
    ProtoMinor int    // 0
    // Header字段用来表示HTTP请求的头域。如果头域（多行键值对格式）为：
    //    accept-encoding: gzip, deflate
    //    Accept-Language: en-us
    //    Connection: keep-alive
    // 则：
    //    Header = map[string][]string{
    //        "Accept-Encoding": {"gzip, deflate"},
    //        "Accept-Language": {"en-us"},
    //        "Connection": {"keep-alive"},
    //    }
    // HTTP规定头域的键名（头名）是大小写敏感的，请求的解析器通过规范化头域的键名来实现这点。
    // 在客户端的请求，可能会被自动添加或重写Header中的特定的头，参见Request.Write方法。
    Header Header
    // Body是请求的主体。
    //
    // 在客户端，如果Body是nil表示该请求没有主体买入GET请求。
    // Client的Transport字段会负责调用Body的Close方法。
    //
    // 在服务端，Body字段总是非nil的；但在没有主体时，读取Body会立刻返回EOF。
    // Server会关闭请求的主体，ServeHTTP处理器不需要关闭Body字段。
    Body io.ReadCloser
    // ContentLength记录相关内容的长度。
    // 如果为-1，表示长度未知，如果>=0，表示可以从Body字段读取ContentLength字节数据。
    // 在客户端，如果Body非nil而该字段为0，表示不知道Body的长度。
    ContentLength int64
    // TransferEncoding按从最外到最里的顺序列出传输编码，空切片表示"identity"编码。
    // 本字段一般会被忽略。当发送或接受请求时，会自动添加或移除"chunked"传输编码。
    TransferEncoding []string
    // Close在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接。
    Close bool
    // 在服务端，Host指定URL会在其上寻找资源的主机。
    // 根据RFC 2616，该值可以是Host头的值，或者URL自身提供的主机名。
    // Host的格式可以是"host:port"。
    //
    // 在客户端，请求的Host字段（可选地）用来重写请求的Host头。
    // 如过该字段为""，Request.Write方法会使用URL字段的Host。
    Host string
    // Form是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据。
    // 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
    Form url.Values
    // PostForm是解析好的POST或PUT的表单数据。
    // 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
    PostForm url.Values
    // MultipartForm是解析好的多部件表单，包括上传的文件。
    // 本字段只有在调用ParseMultipartForm后才有效。
    // 在客户端，会忽略请求中的本字段而使用Body替代。
    MultipartForm *multipart.Form
    // Trailer指定了会在请求主体之后发送的额外的头域。
    //
    // 在服务端，Trailer字段必须初始化为只有trailer键，所有键都对应nil值。
    // （客户端会声明哪些trailer会发送）
    // 在处理器从Body读取时，不能使用本字段。
    // 在从Body的读取返回EOF后，Trailer字段会被更新完毕并包含非nil的值。
    // （如果客户端发送了这些键值对），此时才可以访问本字段。
    //
    // 在客户端，Trail必须初始化为一个包含将要发送的键值对的映射。（值可以是nil或其终值）
    // ContentLength字段必须是0或-1，以启用"chunked"传输编码发送请求。
    // 在开始发送请求后，Trailer可以在读取请求主体期间被修改，
    // 一旦请求主体返回EOF，调用者就不可再修改Trailer。
    //
    // 很少有HTTP客户端、服务端或代理支持HTTP trailer。
    Trailer Header
    // RemoteAddr允许HTTP服务器和其他软件记录该请求的来源地址，一般用于日志。
    // 本字段不是ReadRequest函数填写的，也没有定义格式。
    // 本包的HTTP服务器会在调用处理器之前设置RemoteAddr为"IP:port"格式的地址。
    // 客户端会忽略请求中的RemoteAddr字段。
    RemoteAddr string
    // RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
    // （参见RFC 2616, Section 5.1）
    // 一般应使用URI字段，在客户端设置请求的本字段会导致错误。
    RequestURI string
    // TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息
    // 本字段不是ReadRequest函数填写的。
    // 对启用了TLS的连接，本包的HTTP服务器会在调用处理器之前设置TLS字段，否则将设TLS为nil。
    // 客户端会忽略请求中的TLS字段。
    TLS *tls.ConnectionState
}
```

// 剩余内容

https://www.cnblogs.com/wanghui-garcia/p/10354854.html

## 构建 Restful 服务

符合 REST 约束风格和原则的应用程序或设计就是 RESTful

REST 主要规范了：

1、定位资源的 URL 风格

2、如何对资源操作

REST的主要原则：

1、网络上的所有事物都被抽象为资源

2、每个资源都有唯一的资源标识符

3、同一资源具有多种表现形式（xml，json）

4、对资源的各种操作不会改变资源标识符

5、所有的操作都是无状态的

6、符合 REST 原则的架构方式即可成为 RESTful

**更好的Router**

https://github.com/julienschmidt/httprouter

```go
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name" Hello)

	log.Fatal(http.ListenAndServer(":8080", router))
}
```

**ROA 面向资源的架构（Resource Oriented Architecture）**

# 性能调优

## 性能分析工具
go 内置了很多性能分析工具

**准备工作**

1、安装 graphviz （图形工具）

2、将 $GOPATH/bin 加入 $PATH

3、安装 go-torch

	go get github.com/uber/go-torch

	下载并复制 flamegraph.pl 到 $GOPATH/bin

	将 $GOPATH/bin 添加到 $PATH

**通过文件方式输出 Profile**

1、灵活性高，适用于特定代码端的分析

2、通过手动调用 runtime/pprof 的 API

3、API 相关文档 https://studygolang.com/static/pkgdoc/pkg/runtime_pprof.htm

4、go tool pprof [binary] [binary.prof]

**通过 HTTP 方式输出 Profile**

1、简单，适合于持续性运行的应用

2、在应用程序中导入 import _ "net/http/pprof" 并启动 http server 即可

3、http://<host>:<port>/debug/pprof/

4、go tool pprof http://<host>:<port>/debug/pprof/profile?seconds=10（默认值为30秒）

5、go-torch -seconds 10 http://<host>:<port>/debug/pprof/profile

## 性能调优的示例

**调优过程**

Start -> 设定优化目标 -> 分析系统瓶颈点 -> 优化瓶颈点 -> Start/End

**常见的分析指标**

指标 | 分析
---  | ---
Wall Time | 挂钟时间，程序运行的绝对时间，某个函数运行的绝对时间
CPU Time  | CPU 消耗时间
Block Time | 阻塞时间
Memory allocation | 内存分配
GC times/time spent | GC 次数， GC 耗时

**go 拼接字符串得三种方法**

使用 `+` 拼接字串会严重影响运行性能

1、使用 `bytesBuffer` 拼接字符串

```go
sArr := []string{"a", "b", "c", "d"}
var buffer bytes.Buffer
for i, str := range sArr {
	// Itoa 将整数转换为字符串
	buffer.WriteString(str)
}
fmt.Println(buffer.String)
```
2、构建数组切片得方式拼接子串

```go
sArr := []string{"a", "b", "c", "d"}
fmt.Println(strings.Join(sArr, ""))
```

3、使用 `strings.Builder`
```go
var b strings.Builder
for _, str := range sArr {
	b.WriteString(str)
}
fmt.Println(b.String())
```

**优化思路**

1、生成 cpu.prof 优化耗时时间

2、生成 mem.prof 优化内存

3、生成 goroutine.prof 优化协程

**常见得性能优化点**

- 【CPU】

1、去除不必要的序列化/反序列化：标准的 json 非常消耗性能，可以考虑 easyjson 或者 grpc

2、线程泄漏：goroutine 飞出去之后忘记 stop/close，尤其是 time.ticket 之类的定时器

3、避免不必要的 goroutine，过多的 goroutine 会导致效果过多 CPU

- 【MEM】

1、减少 GC：eg，字符串的传递都是值拷贝，如果数据量打，会差生大量 GC，可以用 byte 数据代替。

2、内存预分配：slice append 的时候空间不够导致不停的 copy 数据来扩大数组大小

- 【DISK】

减少磁盘随机读写IO：磁盘寻道需要花费很多时间

- 【NET】

尽量时候 grpc 代替 http

## 别让性能被锁住

map 不支持线程安全

**sync.Map**

以空间换时间，分为两部分，一是读部分（Read），一是读写部分（Dirty）原子指针指向value，指向相同的value

1、线程安全

2、适合读多写少，且 Key 相对稳定的环境

3、采用空间换时间的方案，并且采用指针的方式间接实现值得映射，所以存储空间会比 build-in map （原生map）大

**Concurrent Map**

适合读写很频繁的情况

原理：普通 map 加锁，锁住真个map，锁的冲突概率很大，而 Concurrent Map，利用分段锁的思想，有多个锁，没把锁锁一段数据，这样在多线程访问不同数据段的锁时，锁竞争的概率比较小。


1、减少锁的影响范围

2、减少发生锁冲突的概率

	sync.Map

	ConcurrentMap

3、避免锁的使用

LAMX Disruptor： https://martinfowler.com/articles/lmax.html

## GC 友好的代码

打开 GC 日志 `GODEBUG=gctrace=1`

**避免内存分配和复制**

复杂对象尽量传递引用

	数组的传递

	结构体的传递

初始化至合适的大小

	自动扩容是有代价的

复用内存

**go tool trace**

```go
//普通程序输出 trace 信息

package main

import (
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()
	// your program here
}


// 测试程序输入 trace 信息
go test -trace trace.out


// 可视化 trace 信息
go tool trace trace.out
```

```go
package auto_growing

import "testing"

const numOfElems = 100000
const times = 1000

func TestAutoGrow(t *testing.T) {
	for i := 0; i < times; i++ {
		s := []int{}
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestProperInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, 100000)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestOverSizeInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, 800000)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{}
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numOfElems)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkOverSizeInit(b *testing.B) {
	for i := 0; i< b.N; i++ {
		s := make([]int, 0, numOfElems * 8)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

// go test -v
=== RUN   TestAutoGrow
--- PASS: TestAutoGrow (0.45s)
=== RUN   TestProperInit
--- PASS: TestProperInit (0.09s)
=== RUN   TestOverSizeInit
--- PASS: TestOverSizeInit (0.40s)
PASS
ok      ch43/gc_friendly/auto_growing   0.946s

// go test -bench=.
goos: linux
goarch: amd64
pkg: ch43/gc_friendly/auto_growing
BenchmarkAutoGrow-2                 2870            432103 ns/op
BenchmarkProperInit-2              12379             99631 ns/op
BenchmarkOverSizeInit-2             2773            401688 ns/op
PASS
ok      ch43/gc_friendly/auto_growing   6.379s
```

# 高可用性服务设计

## 高效字符串拼接

推荐使用 strings.Builder，效率高

```go
package concat_string

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbers = 100

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var builder strings.Builder
		for i := 0; i < numbers; i++ {
			builder.WriteString(strconv.Itoa(i))
		}
		_ = builder.String()
	}
	b.StopTimer()
}

func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++{
		var buf bytes.Buffer
		for i := 0; i < numbers; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
		_ = buf.String()
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}
	}
	b.StopTimer()
}


//
goos: linux
goarch: amd64
pkg: ch44
BenchmarkSprintf-2                 69310             17330 ns/op
BenchmarkStringBuilder-2         1529116               775 ns/op
BenchmarkBytesBuf-2              1236168               987 ns/op
BenchmarkStringAdd-2              224438              4807 ns/op
```

## 面向错误的设计

> Once you accept the failures will heppen, you have the ability to design your system's reaction to the failures

**隔离**

当系统的一部分发生错误时，尽量减少对其他部分的影响，让系统仍能以一定程度的功能进行工作。eg 微内核模式

**隔离错误 - 部署**

微服务

**重用 vs 隔离**

逻辑结构的重用 vs 部署结构的隔离

**冗余**

Load Balancer -> Online Service & Standby Service(备用服务)

Load Balancer -> Online Redundancy(冗余)

**单点失效**

QPS 1500 -> Max QPS 1000 & Max QPS 1000

**限流**

token bucket

Request -> have token? -> response

**慢响应**

不要无休止的等待，给阻塞操作加上一个期限，eg timeout

**错误传递**

断路器配合服务降级。服务出错之后，短路器开启，使用 cache 住的结果或降级服务响应。


## 面向恢复的设计

**健康检查**

- 注意僵尸进程

	池化资源耗尽

	死锁

http ping 检查进程在不在，一定要 cover 关键路径。

**Let it Crash！**

```go
defer func() {
	if err := recover(); err != nil {
		log.Error("recovered panic" err)
	}
}()
```

**构建可恢复的系统**

1、拒绝单体系统

2、面向错误和恢复的设计

	1、在依赖服务不可用时，可以继续存活

	2、快速启动

	3、无状态

**与客户端协商**

服务器：我太忙了，请慢点发送数据

client：好，我一分钟后再发送

## Chaos Engineering 混沌工程

> if something hurts, do it more often!

将故障扼杀再襁褓之中，主动制造故障，测试系统在各种压力下的行为，识别并修复故障问题。

混沌工程以实验发现系统性弱点：

1、定义并测量系统的稳定状态

2、创建假设

3、模拟现实世界中可能发生的事情

4、证明或反驳你的假设

> 目前有如下混沌工程实践方法：模拟数据中心的故障、强制系统时钟不同步、在驱动程序代码中模拟I/O异常、模拟服务之间的延迟、随机引发函数抛异常

- 原则

建立稳定状态的假设

多样化现实世界事件

在生产环境运行实验

持续自动化运行实验

最小化“爆炸半径”

- 相关开源项目

https://github.com/Netflix/chaosmonkey


