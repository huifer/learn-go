package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
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
	getDir(config, "/cron/jobs/")
	//del(config)
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
