package main

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
