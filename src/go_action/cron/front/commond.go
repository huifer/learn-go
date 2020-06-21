package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

var (
	cmd        *exec.Cmd
	err        error
	out        []byte
	resultChan chan *Result
	result     *Result
)

type Result struct {
	err error
	out []byte
}

// 执行一个命令,让他在协程中执行,执行2秒后输出hello
// 执行到1秒的时候杀死
func kill() {
	// 创建结果
	resultChan = make(chan *Result, 1000)

	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
	)
	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		// 命令设置
		cmd = exec.CommandContext(ctx, "bash", "-c", "sleep 2;ls -l;")

		output, err2 := cmd.CombinedOutput()
		// 数据通道
		resultChan <- &Result{
			err: err2,
			out: output,
		}

	}()

	time.Sleep(1 * time.Second)
	// 取消上下文
	cancelFunc()
	// 从队列中获取
	result = <-resultChan

	fmt.Println("结果", string(result.out))
}

// 获取输出结果
func getReturn() {
	// 执行语句设置
	cmd := exec.Command("bash", "-c", "ls -l")
	// 获取输出结果
	out, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return
	}
	// 打印输出
	fmt.Println(string(out))

}
