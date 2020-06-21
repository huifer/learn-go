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
	watchDemo(config)
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
