package main

import "fmt"

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

func main() {
	error03()
}
