# GO mod proxy
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

- go 包管理工具代理

## 代理工具
- 使用 go mod proxy 可以提升包的下载速度, 这里选择的是 [goproxy_cn](https://goproxy.cn/), 具体使用方法详见[官网](https://goproxy.cn/)


~~~
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
~~~




- 编辑器配置

  ![image-20200611081137274](images/image-20200611081137274.png)