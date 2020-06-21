package function

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

/**
闭包
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func adderMain() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func deferFunc() {
	fmt.Println("start ")
	defer fmt.Println("1")
	defer fmt.Println("2")
	// 最后放入位于栈顶优先执行
	defer fmt.Println("3")
	fmt.Println("end ")
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
	adderMain()
	deferFunc()
}
