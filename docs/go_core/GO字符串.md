# GO 字符串
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

## 定义
- 一个Go语言字符串是一个任意字节的常量序列。
## 创建
- 在GO语言中使用双引号`""`或者反引号` 来定义
- 转义字符使用`\`进行描述

```go
package goStrings

import "fmt"

func Create() {
	text1 := "go"
	text2 := `go`

	fmt.Println(text1, text2)
}

```

## 操作
- 获取字符串字节数
```go
text1 := "hello go"
fmt.Println(len(text1))
```
- 字符串切片,使用`[startIndex:endIndex]`规则进行切片
```go
	text1 := "hello go"
	fmt.Println(text1[0])
	fmt.Println(text1[:2])
	fmt.Println(text1[2:])
	fmt.Println(text1[3:5])
```
- 比较
```go
    fmt.Println("go" == "go")
	fmt.Println("go" == "GO")
	
	fmt.Println(strings.Compare("go", "go"))
	fmt.Println(strings.Compare("go", "GO"))

	fmt.Println(strings.EqualFold("go", "go"))
	fmt.Println(strings.EqualFold("go", "GO"))

```
```text
true
false
0
1
true
true
```

说明:
1. `==`区分大小写,编写快,简单 
2. `strings.Compare`区分大小写速度比`==`快
3. `strings.EqualFold`UTF8编码忽略大小写比较


- 字符串拼接,使用`+`进行操作
```go
	text1 := "hello go"
	s := text1 + "123"
```

## strings 包
- 是否存在莫格字符或者字串
```go
	s := "abcdEfg"
	// 只有substr完全在s中才会返回true
	println(strings.Contains(s, "a"))
	// 只要部分chars 在s 中就返回true
	println(strings.ContainsAny(s, "g "))

```

- 字串出现的次数
```go
	s := "abcdEfgab"
	// 字串出现的次数
	print(strings.Count(s, "ab"))
```

- 字符串分割
    - Split 切割 和 SplitN 切割区别在N,N为切成几份,只切到第N个切割符号
```go
	// 字符串分割
	s2 := "a,b,c,d"
	split := strings.Split(s2, ",")
	for _, s3 := range split {
		println(s3)
	}
	println("----")
	n := strings.SplitN(s2, ",", 2)
	for _, s3 := range n {
		println(s3)
	}
```

- 是否存在前后缀
```go
	// 前缀
	println(strings.HasPrefix(s2, "a"))
	// 后缀
	println(strings.HasSuffix(s2, "a"))
```
- 字符串替换
```go
// 字符串替换
println(strings.Replace(s2, "a", "abc", -1))
```