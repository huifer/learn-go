package utils

import (
	"fmt"
	"time"
)

func PrintTime() {
	t := time.Now()
	fmt.Println("当前时间: ", t)

}
