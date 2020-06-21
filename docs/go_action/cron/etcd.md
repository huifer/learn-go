# ETCD
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

## 核心特性
1. 将数据存储在集群中的高可用 K-V 存储
1. 允许应用实时监听存储中的 K-V 变化
1. 能够容忍单点故障,能够应对网络分区
1. key 有序存储


## 环境搭建
- 二进制包: https://github.com/etcd-io/etcd/releases
    - 通过 github 进行搜索 etcd 下载二进制包
    - 选择你的开发环境进行安装即可
    
    
```shell script
unzip etcd-v3.4.9-darwin-amd64.zip
cd etcd-v3.4.9-darwin-amd64
```

- 启动,开放任意网卡

```shell script
nohup ./etcd --listen-client-urls 'http://0.0.0.0:2379' --advertise-client-urls 'http://0.0.0.0:2379' &
ps -ef | grep etcd 
```

```text
501 20863   990   0  5:09下午 ttys000    0:00.12 ./etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379
501 21260   990   0  5:09下午 ttys000    0:00.00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn etcd
```

- 日志查看

```shell script
less nohup.out
```

```text
2020-06-21 17:09:06.714982 I | etcdserver: published {Name:default ClientURLs:[http://0.0.0.0:2379]} to cluster cdf818194e3a8c32
2020-06-21 17:09:06.715006 I | embed: ready to serve client requests
2020-06-21 17:09:06.715015 I | etcdserver/api: enabled capabilities for version 3.4
2020-06-21 17:09:06.716671 N | embed: serving insecure client requests on [::]:2379, this is strongly discouraged!
```

- 启动成功


## etcdctl 命令行


```text
NAME:
	etcdctl - A simple command line client for etcd3.

USAGE:
	etcdctl [flags]

VERSION:
	3.4.9

API VERSION:
	3.4


COMMANDS:
	alarm disarm		Disarms all alarms
	alarm list		Lists all alarms
	auth disable		Disables authentication
	auth enable		Enables authentication
	check datascale		Check the memory usage of holding data for different workloads on a given server endpoint.
	check perf		Check the performance of the etcd cluster
	compaction		Compacts the event history in etcd
	defrag			Defragments the storage of the etcd members with given endpoints
	del			Removes the specified key or range of keys [key, range_end)
	elect			Observes and participates in leader election
	endpoint hashkv		Prints the KV history hash for each endpoint in --endpoints
	endpoint health		Checks the healthiness of endpoints specified in `--endpoints` flag
	endpoint status		Prints out the status of endpoints specified in `--endpoints` flag
	get			Gets the key or a range of keys
	help			Help about any command
	lease grant		Creates leases
	lease keep-alive	Keeps leases alive (renew)
	lease list		List all active leases
	lease revoke		Revokes leases
	lease timetolive	Get lease information
	lock			Acquires a named lock
	make-mirror		Makes a mirror at the destination etcd cluster
	member add		Adds a member into the cluster
	member list		Lists all members in the cluster
	member promote		Promotes a non-voting member in the cluster
	member remove		Removes a member from the cluster
	member update		Updates a member in the cluster
	migrate			Migrates keys in a v2 store to a mvcc store
	move-leader		Transfers leadership to another etcd cluster member.
	put			Puts the given key into the store
	role add		Adds a new role
	role delete		Deletes a role
	role get		Gets detailed information of a role
	role grant-permission	Grants a key to a role
	role list		Lists all roles
	role revoke-permission	Revokes a key from a role
	snapshot restore	Restores an etcd member snapshot to an etcd directory
	snapshot save		Stores an etcd node backend snapshot to a given file
	snapshot status		Gets backend snapshot status of a given file
	txn			Txn processes all the requests in one transaction
	user add		Adds a new user
	user delete		Deletes a user
	user get		Gets detailed information of a user
	user grant-role		Grants a role to a user
	user list		Lists all users
	user passwd		Changes password of user
	user revoke-role	Revokes a role from a user
	version			Prints the version of etcdctl
	watch			Watches events stream on keys or prefixes

OPTIONS:
      --cacert=""				verify certificates of TLS-enabled secure servers using this CA bundle
      --cert=""					identify secure client using this TLS certificate file
      --command-timeout=5s			timeout for short running command (excluding dial timeout)
      --debug[=false]				enable client-side debug logging
      --dial-timeout=2s				dial timeout for client connections
  -d, --discovery-srv=""			domain name to query for SRV records describing cluster endpoints
      --discovery-srv-name=""			service name to query when using DNS discovery
      --endpoints=[127.0.0.1:2379]		gRPC endpoints
  -h, --help[=false]				help for etcdctl
      --hex[=false]				print byte strings as hex encoded strings
      --insecure-discovery[=true]		accept insecure SRV records describing cluster endpoints
      --insecure-skip-tls-verify[=false]	skip server certificate verification
      --insecure-transport[=true]		disable transport security for client connections
      --keepalive-time=2s			keepalive time for client connections
      --keepalive-timeout=6s			keepalive timeout for client connections
      --key=""					identify secure client using this TLS key file
      --password=""				password for authentication (if this option is used, --user option shouldn't include password)
      --user=""					username[:password] for authentication (prompt if password is not supplied)
  -w, --write-out="simple"			set the output format (fields, json, protobuf, simple, table)

```

- 设置 `./etcdctl put "name" "huifer"`
- 获取 `./etcdctl get "name" `
- 删除 `./etcdctl del "name"`



### get 命令帮助

```shell script
 ./etcdctl get -h
```


```text
NAME:
	get - Gets the key or a range of keys

USAGE:
	etcdctl get [options] <key> [range_end] [flags]

OPTIONS:
      --consistency="l"			Linearizable(l) or Serializable(s)
      --from-key[=false]		Get keys that are greater than or equal to the given key using byte compare
  -h, --help[=false]			help for get
      --keys-only[=false]		Get only the keys
      --limit=0				Maximum number of results
      --order=""			Order of results; ASCEND or DESCEND (ASCEND by default)
      --prefix[=false]			Get keys with matching prefix
      --print-value-only[=false]	Only write values when using the "simple" output format
      --rev=0				Specify the kv revision
      --sort-by=""			Sort target; CREATE, KEY, MODIFY, VALUE, or VERSION

GLOBAL OPTIONS:
      --cacert=""				verify certificates of TLS-enabled secure servers using this CA bundle
      --cert=""					identify secure client using this TLS certificate file
      --command-timeout=5s			timeout for short running command (excluding dial timeout)
      --debug[=false]				enable client-side debug logging
      --dial-timeout=2s				dial timeout for client connections
  -d, --discovery-srv=""			domain name to query for SRV records describing cluster endpoints
      --discovery-srv-name=""			service name to query when using DNS discovery
      --endpoints=[127.0.0.1:2379]		gRPC endpoints
      --hex[=false]				print byte strings as hex encoded strings
      --insecure-discovery[=true]		accept insecure SRV records describing cluster endpoints
      --insecure-skip-tls-verify[=false]	skip server certificate verification
      --insecure-transport[=true]		disable transport security for client connections
      --keepalive-time=2s			keepalive time for client connections
      --keepalive-timeout=6s			keepalive timeout for client connections
      --key=""					identify secure client using this TLS key file
      --password=""				password for authentication (if this option is used, --user option shouldn't include password)
      --user=""					username[:password] for authentication (prompt if password is not supplied)
  -w, --write-out="simple"			set the output format (fields, json, protobuf, simple, table)
```

- 前缀查询 `--prefix`




## GO 使用 etcd
- https://github.com/etcd-io/etcd/tree/master/clientv3

### 安装
`go get go.etcd.io/etcd/clientv3`


### go 客户端使用
#### 创建客户端


```go
config := clientv3.Config{
		// 链接地址
		Endpoints:   []string{"0.0.0.0:2379"},
		// 超时时间
		DialTimeout: 5 * time.Second,
}

```


#### 设置数据

```go
func putDemo(config clientv3.Config) {
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
	}
	kv := clientv3.NewKV(client)

	put, err := kv.Put(context.TODO(), "/cron/jobs/job1", "hello2", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(put)

	if put.PrevKv != nil {
		fmt.Println("覆盖了原始数据", string(put.PrevKv.Value))
	}
}

```

- 命令行获取

```shell script
./etcdctl get "/cron/" --prefix                                                                                            ✔  ⚙  3204  18:03:10


/cron/jobs/job1
hello2
```


#### 获取

```go
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

```

- 输出结果

```text
&{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:6 raft_term:2  [key:"/cron/jobs/job1" create_revision:4 mod_revision:6 version:3 value:"hello2" ] false 1 {} [] 0}
```

#### 删除

```go

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

```

- 执行结果

```text
&{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:8 raft_term:2  1 [] {} [] 0}

```


#### 监听