# GO 接口

## 定义
- 接口是一组方法的集合

## 声明
- 关键字 `interface`

```go
type InterfaceName interface{
    Method1
    Method2
}
```

- 接口的命名一般以 **`er`**结尾
- 接口内部的方法不需要 `func`
- 接口内部的方法没有实现,只有声明


## 使用
- 定义一个输出接口

```go
type Printer interface {
	Info(s string)
	Error(s string)
}

```

1. 接口没有实现直接调用会报错


```go
func main() {
	var i Printer

	i.Error("hello")

}

```

- 错误信息

```text
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x0 pc=0x45acc6]

goroutine 1 [running]:
main.main()
	C:/Users/admin/GolandProjects/learn-go/src/go_core/goInterface/GoInterface.go:11 +0x26

```

- 实现接口

```go
import "fmt"

type Printer interface {
	Info(s string)
	Error(s string)
}


type Logger struct{}

func (Logger) Info(s string) {
	fmt.Println("info: ", s)

}

func (Logger) Error(s string) {
	fmt.Println("error: ", s)
}
```

## 类型断言
- 语法 : `i.(TypeName)`
    - 如果接口`i`是 TypeName 不会报错

```go
	ca1 := i.(Logger)
	fmt.Println(ca1)
    // 报错
	ca2 := i.(string)
	fmt.Println(ca2)
```

- 类型判断失败会抛出异常

```text
.\GoInterface.go:30:10: impossible type assertion:
	string does not implement Printer (missing Error method)
```

- 类型判断使用

```go
package main

import "fmt"

type Printer interface {
	Info(s string)
	Error(s string)
}

type Logger struct{}

func (Logger) Info(s string) {
	fmt.Println("info: ", s)

}

func (Logger) Error(s string) {
	fmt.Println("error: ", s)
}

type Oc interface {
	Ic(s string)
}

type OcImp struct {

}

func (OcImp) Ic(s string) {
	fmt.Println("oc", s)

}

func main() {
	var i Printer

	i = Logger{}


	if o, ok := i.(Logger); ok {
		o.Info("接口")
	}
	if o, ok := i.(Oc); ok {
		o.Ic("okokokok")
	}
}

```

