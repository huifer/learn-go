package basicGrammar

import (
	"fmt"
	"math"
)

/**
int类型
*/
func intValue() {
	var aNumb int32 = 123
	var bNumb int = int(int8(aNumb))
	fmt.Print(bNumb)
}

/**
float 类型
*/
func floatValue() {
	var cFloat float32 = math.Pi
	fmt.Print(int(cFloat))
}

/**
字符串类型
*/
func strValue() {
	var st string = "hello go"
	fmt.Print(st)
}

/**
字符串拼接
*/
func strAdd() {
	var st1 = "hello"
	var st2 = "go"
	var st3 = st1 + " " + st2
	fmt.Print(st3)

}

/**
boolean 类型
*/
func booleanValue() {
	var b = true
	fmt.Print(b)
}

/**
常量
*/
const (
	lang = "go"
)

/**
常量获取
*/
func constValue() {
	fmt.Print(lang)
}

/**
两个数相加
*/
func addTwoNumb(a int, b int) {
	var c = a + b
	fmt.Print(c)
}

/**
流程控制
*/
func ifValue() {
	var a = 1
	if a > 0 {
		fmt.Print(">0")
	} else if a > 1 {
		fmt.Print(">1")
	} else {
		fmt.Print("其他")
	}
}

/**
switch
*/
func switchValue() {
	var sw = "a"
	switch sw {
	case "a", "c":
		fmt.Print(sw)
	case lang:
		fmt.Print(lang)
	default:
		fmt.Print("default")
	}
}

/**
for 循环
*/
func forValue() {
	for i := 0; i < 5; i++ {
		fmt.Print("hello world")
	}
}

/**
for 循环
*/
func forValue2() {
	var i = []string{"go", "java", "python"}
	for index, value := range i {
		fmt.Print(index, value)
	}
}

func gotoValue() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			// 强制转到
			goto gotoHere
		}
	}

gotoHere:
	fmt.Print("in goto ")
}

func hello(a int, b string) string {
	return "zs"
}

func swap(x, y string) (string, string) {
	return y, x
}

/**
无参闭包
*/
func add() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func add2(x, y int) func() (int, int) {
	i := 0
	return func() (x int, y int) {
		i++
		return i, x + y
	}
}
