package main

import "fmt"

func add2num(x, y int) int {
	return x + y
}
func swap2(x, y int) (int, int) {
	return y, x
}

func returnName() (a int, b int) {
	a = 1
	b = 2
	return a, b
}

func printHello() {
	fmt.Println("hello world")
}

func printHello2(string2 string) {
	fmt.Println("hello" + string2)
}

func callBackMain(i int, f func(int)) {
	f(i)
}

func callBackFunc(i int) {
	fmt.Println("匿名函数回调", i)
}

func main() {
	i := add2num(1, 2)
	fmt.Print(i)
	fmt.Print("\n")
	i2, i3 := swap2(1, 2)
	fmt.Print(i2)
	fmt.Print(i3)
	fmt.Print("\n")
	c, ad := returnName()
	fmt.Println(c, ad)

	var funcValue func()
	funcValue = printHello
	funcValue()

	var f = printHello2
	f("go")

	func(i int) {
		fmt.Println(i)
	}(1000)

	callBackMain(100, callBackFunc)
}
