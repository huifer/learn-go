# GO基础语法
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go
>


## 关键字列表 
1. break
1. default
1. func
1. interface
1. select 
1. case
1. defer 
1. go
1. map
1. struct
1. chan 
1. eles
1. goto
1. package 
1. switch
1. const
1. fallthrough
1. if 
1. range
1. type
1. continue 
1. for 
1. import 
1. return 
1. var 


### 引导程序整体结构的关键字
1. package : 定义包名的关键字
1. import : 导入报名关键字
1. const : 声明倡廉关键字
1. var : 声明变量关键字
1. func  : 函数定义关键字
1. defer : 延迟执行关键字
1. go : 并发语法糖关键字
1. return : 函数返回关键字

### 声明复合数据结构的关键字
1. struct : 定义结构类型关键字
1. interface : 定义接口类型关键字
1. map : 声明或创建map类型关键字
1. chan : 声明或创建通道类型关键字

### 控制程序结构的关键字
1. if \ else : 条件分支语句关键字
1. for range break continue : for 循环关键字
1. switch select type case default fallthrough : switch 和 select 语句关键字
1. goto : 跳转语句关键字 

## 数据类型
### 整数
1. byte
1. int
1. int8
1. int16
1. int32
1. int64
1. uint
1. uint8
1. uint16
1. uint32
1. uint64
1. uintptr
### 浮点数
1. float32
1. float64


### 复数
1. complex64
1. complex128

### 字符和字符串
1. string
1. rune

### 接口
1. error

### 布尔
1. bool

### 常量值
1. true false : bool类型, 真和假
1. iota : 用在连续的枚举类型声明中
1. nil : 指针、应用类型的变量默认值

### 空白标识符
1. `_` : 空白标识符用来声明一个匿名的变量.

## 注释

- 单行注释：`//`
- 多行注释：`/***/`

## 标识符

- 标识符是用户或者系统定义的有含义单词的组合.
- 标识符一般以**字母或者下划线**开头，**大小写敏感**



## 常量

- 常量使用**`const`**修饰，表示**只读，不能修改**.

```go
const (
	lang string = "go"
)
```





## 变量

- 变量使用**`var`**修饰



## 关键字

- GO语言关键字表如下

| break    | default    | func   | interface | select |
| -------- | ---------- | ------ | --------- | ------ |
| case     | defer      | go     | map       | struct |
| chan     | else       | goto   | package   | switch |
| const    | fallthough | if     | range     | type   |
| continue | for        | import | return    | var    |



## 操作符

数字操作符：+, -, *, /, %

比较运算符：>, >=, <, <=, ==, !=





## 数据类型

- 数字类型：int8, int16, int32, int64, uint8, uint16, uint32, uint64
- bool 类型：ture, false
- 浮点类型：float32, float64
- 字符类型：byte
- 字符串类型：string



## 数据类型代码
- [GO语言基础语法](/basicGrammar/GoBasicGrammar.go)

```java
package main

import (
	"fmt"
	"math"
)

/**
int类型
*/
func intValue() {
	var aNumb int32 = 123
	var bNumb int = int(int8(aNumb))
	fmt.Print(bNumb)
}

/**
float 类型
*/
func floatValue() {
	var cFloat float32 = math.Pi
	fmt.Print(int(cFloat))
}

/**
字符串类型
*/
func strValue() {
	var st string = "hello go"
	fmt.Print(st)
}

/**
字符串拼接
*/
func strAdd() {
	var st1 = "hello"
	var st2 = "go"
	var st3 = st1 + " " + st2
	fmt.Print(st3)

}

/**
boolean 类型
*/
func booleanValue() {
	var b = true
	fmt.Print(b)
}

/**
常量
*/
const (
	lang = "go"
)

/**
常量获取
*/
func constValue() {
	fmt.Print(lang)
}

/**
两个数相加
 */
func addTwoNumb(a int, b int) {
	var c = a + b
	fmt.Print(c)
}


```



## 流程控制

### 分支判断

- `if else`

```go
func ifValue(){
var a = 1
if a > 0 {
    fmt.Print(">0")
} else if a >1 {
    fmt.Print(">1")
}else {
    fmt.Print("其他")
}
}
```

  

- `switch case `

```go
func switchValue() {
var sw = "a"
switch sw {
case "a", "c":
    fmt.Print(sw)
case lang:
    fmt.Print(lang)
default:
    fmt.Print("default")
}
}

```

  



## 循环

- for

  ```go
  func forValue() {
  	for i := 0; i < 5; i++ {
  		fmt.Print("hello world")
  	}
  }
  ```

  

## goto

- 强制转到

```go
func gotoValue() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			// 强制转到
			goto gotoHere
		}
	}

gotoHere:
	fmt.Print("in goto ")
}
```



## 函数

### 函数定义

- 函数定义使用

  ```go
  func 函数名称( [参数列表,变量名 类型] ) [返回值类型] {
     //函数体
  }
  ```

  ```go
  func hello(a int, b string) string {
  	return "zs"
  }
  
  ```

### 多返回值

```go
func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("go", "java")
	fmt.Print(a, b)
}
```



### 匿名函数

- 匿名函数没有函数名，只有函数体，可以直接被当做一种类型赋值给函数类型的变量，匿名函数也往往以变量的方式被传递，匿名函数经常被用于实现回调函数、闭包

  ```go
  func add() func() int {
  	i := 0
  	return func() int {
  		i += 1
  		return i
  	}
  }
  ```

  

## 练习

1. 输出乘法表
2. 对于数字n求阶乘
3. 对字符串进行统计英文，空格数字有多少个

- [答案](/Answer/BasicGrammarAnswer.go)