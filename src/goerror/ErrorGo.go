package goerror

import (
	"errors"
	"fmt"
)

func error01(x, y int) (int, error) {
	if x == y {
		return 0, errors.New("两数相等")
	} else {
		return x + y, nil

	}
}

func error02(x int) {
	if x == 0 {
		errors.New("x等于0")
	}
}

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

func error3(name string) (err error) {
	return errors.New(name)
}
