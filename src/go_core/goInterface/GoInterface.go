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

func main() {
	var i Printer

	i = Logger{}
	i.Error("hello")
	i.Info("hello")

}
