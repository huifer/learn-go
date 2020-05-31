# GO 异常
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go
>
## 异常定义
- Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
- error 为接口

```go
type error interface {
	Error() string
}
```

## 异常的使用

```go
func error02(x int) {
	if x == 0 {
		errors.New("x等于0")
	}
}
```

- go 没有 `try{}catch{}finally{}` go 使用 `defer panic recover`处理

```go
func error03() {
	//捕获 + 处理
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error : ", err)
		}
	}()
	a := 1
	b := 0
	res := a / b
	fmt.Println(res)
}

```

## 自定义异常

```go
func error3(name string)(err error){
	return errors.New(name)
}
```