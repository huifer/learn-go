# GO 异常
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go
>
## 异常定义
- Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
- error 为接口
- 在读个返回值的函数中, 通常 error 作为最后一个返回值
- 如果函数返回值为 error 类型, 请先使用 `error != nil`进行判断
- defer 语句应该放在 error 的判断后面


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


## panic 
- 触发 panic 的两种情况
    1. 主动调用 panic 函数
    1. 程序运行时出现的异常,go捕捉到了在处理

## recover
- recover 用来捕获 panic , 阻止 panic 继续向上抛出异常