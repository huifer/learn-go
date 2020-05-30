# GO 函数
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go
>
## 定义
- [前文](/doc/go_核心编程/02-GO基础语法.mdO基础语法.md)简单描述过函数如何定义以及使用,这里提出重新讲解.
- 函数的祖成包括: 函数名、返回值列表、参数列表(形参).
- 基本函数定义如: 
```go
func 函数名(参数列表) (返回值列表){
    // 函数体
}
```
## 参数列表
```go
func add2num(x, y int) int {
	return x + y
}

func main() {
	i := add2num(1, 2)
	fmt.Print(i)
}

```

- 在上面代码中`add2num`的参数列表为x、y类型是int,`main`没有参数列表

## 返回值列表
- go语言的返回值可以定义多个.
```go
func swap2(x, y int) (int, int) {
	return y, x
}
```
- 上述代码中传入x和y,返回值为xy交换后的值.注意通过`(int ,int)`控制返回值类型


### 带有变量名的返回值
```GO
func returnName()(a int ,b int  ){
	a = 1
	b = 2
	return a , b
}

func main(){
	c, ad := returnName()
    fmt.Println(c , ad)
}
```
- 上述代码将a,b作为返回值,目前个人不知道这段代码可以在什么场景下使用T_T
    - 疑惑: 返回值定义名字有什么意义? 函数的运行结果变量名本就可以自己取想取什么都可.
    
    
## 调用函数
`返回值变量列表  := 函数名(形式参数列表)`


## 函数传递
- GO语言中函数也是变量(有点JavaScript的意思)
### 无参函数
```go

func printHello(){
	fmt.Println("hello world")
}

func main(){
    var funcValue func()
	funcValue = printHello
	funcValue()
}
```
### 有参函数
```go
func printHello2(string2 string) {
	fmt.Println("hello" + string2)
}

func main(){
	var f = printHello2
	f("go")
}
```

## 匿名函数
- 匿名函数顾名思义,没有名字的函数就是匿名函数,匿名函数不可单独定义,需要再函数中定义
### 定义匿名函数
```go
func (参数列表) (返回值列表){
    // 函数体
}
```
```go
	func(i int){
		fmt.Println(i)
	}(1000)
```
### 回调函数
- 匿名函数可用作回调函数使用(更加像JS了...)
```go
func callBackMain(i int, f func(int)) {
	f(i)
}

func callBackFunc(i int) {
	fmt.Println("匿名函数回调", i)
}
func main(){
	callBackMain(100, callBackFunc)
}
```


## 闭包
- 闭包的体现形式，能常就是用函数返回另一个函数
- **闭包的概念** ：是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环境（作用域）。
- **闭包的价值** : 闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。
- **使用闭包的注意点** :由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包
- Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在。
```go

/**
闭包
 */
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func adderMain() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

## 延迟执行 defer 
- Go语言的 `defer` 语句会将其后面跟随的语句进行延迟处理，在 `defer` 归属的函数即将返回时，将延迟处理的语句按 `defer` 的逆序进行执行，也就是说，先被 `defer` 的语句最后被执行，最后被 `defer` 的语句，最先被执行。
```go
func deferFunc() {
	fmt.Println("start ")
	defer fmt.Println("1")
	defer fmt.Println("2")
	// 最后放入位于栈顶优先执行
	defer fmt.Println("3")
	fmt.Println("end ")
}

```
- 输出结果: 
```text
start 
end 
3
2
1
```


## 测试函数
- Go语言自带了 testing 测试包，可以进行自动化的单元测试，输出结果验证，并且可以测试性能。
### 测试规则
1. 准备源码文件
2. 命名以`_test.go`结尾,测试函数以`Test`作为方法前缀
### 单元测试
```go
package function

func add4(x, y int) int {
	return x + y
}

```
- 测试用例
```go
package function

import "testing"

func Test_add4(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 添加单元测试数据
		{name: "1+1", args: args{1, 1}, want: 2,},
		{name: "1+2", args: args{1, 2}, want: 2,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add4(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add4() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
- 测试结果
```
=== RUN   Test_add4
=== RUN   Test_add4/1+1
=== RUN   Test_add4/1+2
--- FAIL: Test_add4 (0.00s)
    --- PASS: Test_add4/1+1 (0.00s)
    --- FAIL: Test_add4/1+2 (0.00s)
        functionTest_test.go:23: add4() = 3, want 2
FAIL
```

### 性能测试
- 命名规则: 以 `Benchmark`作为前缀
```go
func BenchmarkAdd(t *testing.B){
	for i := 0; i < t.N; i++ {
		add4(1, 2)
	}
}

```
- 运行结果
```text
goos: windows
goarch: amd64
pkg: lern-go/src/function
BenchmarkAdd-8   	1000000000	         0.510 ns/op
PASS

```
- 结果说明: 执行 1000000000 次 消费 0.510 纳秒
- 命令行模式下使用`go test -test.bench='.*'`进行单元测试
### 覆盖率测试
- 命令行:`go test -cover`
```text
PASS
coverage: 2.4% of statements
ok      lern-go/src/function    0.338s
```

## 练习
1. 创建函数大于5输出true,其他输出false
2. 对1进行测试
- [答案](/Answer/FunctionAnswer.go)