package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

func main() {

	config := clientv3.Config{
		// 链接地址
		Endpoints: []string{"0.0.0.0:2379"},
		// 超时时间
		DialTimeout: 5 * time.Second,
	}
	//putDemo(config)
	//getDir(config, "/cron/jobs/")
	//del(config)
	OptimisticLock(config)
}

// 乐观锁
func OptimisticLock(config clientv3.Config) {
	var (
		leaseGrantResponse *clientv3.LeaseGrantResponse
		err                error
		keepResponse       *clientv3.LeaseKeepAliveResponse
		ctx                context.Context
		cancelFunc         context.CancelFunc
	)
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	//1. 上锁(创建租约)
	lease := clientv3.NewLease(client)

	// 创建十秒的租约
	if leaseGrantResponse, err = lease.Grant(context.TODO(), 3); err != nil {
		fmt.Println(err)
		return
	}

	leaseId := leaseGrantResponse.ID

	// 取消租约的上下文
	ctx, cancelFunc = context.WithCancel(context.TODO())
	// 函数结束后,停止续租
	defer cancelFunc()
	// 主动释放,立即的
	defer lease.Revoke(context.TODO(), leaseId)

	alive, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			select {
			case keepResponse = <-alive:

				if keepResponse == nil {
					fmt.Println("租约失效")
					goto END

				} else {
					fmt.Println("收到自动续约应答", keepResponse.ID)
				}
			}
		END:
		}
	}()

	kv := clientv3.NewKV(client)
	// 事务
	txn := kv.Txn(context.TODO())
	// if key 不存在 then 设置 else 获取所失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job"), "=", 0)).Then(clientv3.OpPut("/cron/lock/job", "ddd", clientv3.WithLease(leaseId))).Else(clientv3.OpGet("/cron/lock/job"))
	commit, err := txn.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 判断是否抢到
	if !commit.Succeeded {
		fmt.Println("锁被占用", string(commit.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	//2. 业务
	fmt.Println("处理业务")
	time.Sleep(5 * time.Second)
	// 3. 释放锁(取消自动续租,释放租约)

}

func watchDemo(config clientv3.Config) {
	var (
		watchResp clientv3.WatchResponse
		event     *clientv3.Event
	)
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
	}
	kv := clientv3.NewKV(client)

	// 变化
	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job3", "value1")
			kv.Delete(context.TODO(), "/cron/jobs/job3")
			time.Sleep(1 * time.Second)
		}
	}()

	// 获取
	get, err := kv.Get(context.TODO(), "/cron/jobs/job3")
	if err != nil {
		fmt.Println(err)
	}
	if len(get.Kvs) != 0 {
		fmt.Println("当前值存在", string(get.Kvs[0].Value))

	}

	watchStartRevision := get.Header.Revision + 1

	// 创建一个watcher(监听器)
	watcher := clientv3.NewWatcher(client)
	fmt.Println("从该版本向后监听:", watchStartRevision)
	ctx, cancelFunc := context.WithCancel(context.TODO())

	//5秒钟后取消
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})

	// 将监听的数据网队列中放置
	watchRespChan := watcher.Watch(ctx, "/cron/jobs/job3", clientv3.WithRev(watchStartRevision))

	// 处理kv变化事件
	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			// 不同的处理事件进行不同的处理
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}
}

func del(config clientv3.Config) {
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
	}
	kv := clientv3.NewKV(client)
	response, err := kv.Delete(context.TODO(), "/cron/jobs/job2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
func getDir(config clientv3.Config, pre string) {
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
	}
	kv := clientv3.NewKV(client)

	get, err := kv.Get(context.TODO(), pre, clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(get.Kvs)
	for _, keyValue := range get.Kvs {

		fmt.Println(string(keyValue.Key), string(keyValue.Value))
	}

}

func getDemo(config clientv3.Config) {
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)

	}
	kv := clientv3.NewKV(client)

	if get, err := kv.Get(context.TODO(), "/cron/jobs/job1"); err == nil {
		if get.Count > 0 {
			fmt.Println(get)
		}

	}

}

func putDemo(config clientv3.Config) {
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
	}
	kv := clientv3.NewKV(client)

	put, err := kv.Put(context.TODO(), "/cron/jobs/job2", "hello2", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(put)

	if put.PrevKv != nil {
		fmt.Println("覆盖了原始数据", string(put.PrevKv.Value))
	}
}
