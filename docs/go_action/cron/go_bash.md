# GO bash
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

## 获取 bash 执行结果

```go
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
```

- 执行结果

```text
total 136
-rw-r--r--  1 huifer  staff  11558 May 31 13:52 LICENSE
-rw-r--r--  1 huifer  staff   1455 Jun 21 14:12 README.md
-rw-r--r--  1 huifer  staff    290 May 31 17:59 _404.md
-rw-r--r--  1 huifer  staff     26 May 31 13:52 _config.yml
-rw-r--r--  1 huifer  staff     64 May 31 19:18 _coverpage.md
drwxr-xr-x  3 huifer  staff     96 May 31 13:52 bin
drwxr-xr-x  3 huifer  staff     96 Jun 21 14:12 doc
drwxr-xr-x  7 huifer  staff    224 Jun  5 19:05 docs
-rw-r--r--  1 huifer  staff   1684 Jun 20 19:00 go.mod
-rw-r--r--  1 huifer  staff  24016 Jun 20 19:04 go.sum
-rw-r--r--  1 huifer  staff   2295 May 31 19:59 index.html
drwxr-xr-x  3 huifer  staff     96 May 31 13:52 pkg
drwxr-xr-x  8 huifer  staff    256 Jun 20 19:07 src
-rw-r--r--  1 huifer  staff    966 Jun 20 19:00 summary.md
-rw-r--r--  1 huifer  staff    351 May 31 13:52 开发方向.md

```

## 杀死进程


```go

var (
	cmd        *exec.Cmd
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

```

