package goStrings

import (
	"fmt"
	"strings"
)

func Create() {
	text1 := "go"
	text2 := `go`

	fmt.Println(text1, text2)
}

func Operation() {
	text1 := "hello go"
	// 字节数
	fmt.Println(len(text1))
	// 切片
	fmt.Println(text1[0])
	fmt.Println(text1[:2])
	fmt.Println(text1[2:])
	fmt.Println(text1[3:5])

	// 字符串比较
	fmt.Println("go" == "go")
	fmt.Println("go" == "GO")

	fmt.Println(strings.Compare("go", "go"))
	fmt.Println(strings.Compare("go", "GO"))

	fmt.Println(strings.EqualFold("go", "go"))
	fmt.Println(strings.EqualFold("go", "GO"))

	// 字符串拼接
	s := text1 + "123"
	fmt.Println()
}
