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
	fmt.Println(s)

}

func UseStrings() {
	s := "abcdEfgab"
	// 只有substr完全在s中才会返回true
	println(strings.Contains(s, "a"))
	// 只要部分chars 在s 中就返回true
	println(strings.ContainsAny(s, "g "))

	// 字串出现的次数
	println(strings.Count(s, "ab"))

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
	println("----")
	// 前缀
	println(strings.HasPrefix(s2, "a"))
	// 后缀
	println(strings.HasSuffix(s2, "a"))
	// 字符串替换
	println(strings.Replace(s2, "a", "abc", -1))

}
