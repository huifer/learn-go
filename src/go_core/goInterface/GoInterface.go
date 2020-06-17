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
	i.Error("hello")
	i.Info("hello")

	ca1 := i.(Logger)
	fmt.Println(ca1)
	// 类型判断失败
	//ca2 := i.(string)
	//fmt.Println(ca2)

	if o, ok := i.(Logger); ok {
		o.Info("接口")
	}
	if o, ok := i.(Oc); ok {
		// 类型不匹配则进入不了下面这段代码
		o.Ic("okokokok")
	}
}
