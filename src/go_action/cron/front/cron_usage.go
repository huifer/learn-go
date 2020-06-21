package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	cronTable()
}

// 任务结构
type CronJob struct {
	// cron 表达式
	expr *cronexpr.Expression
	// 下一次的执行时间
	nextTime time.Time
}

// 多任务执行
func cronTable() {

	var (
		cronJob *CronJob
		expr    *cronexpr.Expression
		// 调度表
		scheduleTable map[string]*CronJob
	)
	scheduleTable = make(map[string]*CronJob)
	// 任务定义
	// 每五秒执行
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(time.Now()),
	}
	scheduleTable["5s"] = cronJob
	// 每十秒执行
	expr = cronexpr.MustParse("*/10 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(time.Now()),
	}
	scheduleTable["10s"] = cronJob

	// 协程定时检查调度表
	go func() {
		var (
			jobName string
			cronJob *CronJob
		)
		for {
			for jobName, cronJob = range scheduleTable {

				// 是否过期
				if cronJob.nextTime.Before(time.Now()) || cronJob.nextTime.Equal(time.Now()) {
					// 执行
					go func(jobName string) {
						fmt.Println("当前执行的是", jobName)
					}(jobName)
					// 计算下一次执行时间
					fmt.Println("设置下一次执行时间")
					cronJob.nextTime = cronJob.expr.Next(time.Now())
				}
			}
			select {
			// 等100毫秒在执行循环
			case <-time.NewTimer(100 * time.Millisecond).C:

			}
		}
	}()

	time.Sleep(300 * time.Second)
}

func cron() {

	cron := "* * * * *"
	expr, err := cronexpr.Parse(cron)
	if err != nil {
		fmt.Println(err)
	}
	//输入当前时间 得到下一次执行时间
	now := time.Now()
	next := expr.Next(now)
	fmt.Println("当前时间 ", now, "下次执行时间 ", next)

	// 创建一个定时器,等待定时器超时以后做一个任务
	time.AfterFunc(next.Sub(now), func() {
		fmt.Println("hello")
	})
	time.Sleep(50 * time.Second)

}
