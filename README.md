# Goods Srv

This is the Goods service with fqdn go.micro.srv.v1.goods.

## Getting Started
### Prerequisites

> [Micro](https://github.com/micro/micro)

## Install Micro

```shell
go get -u github.com/micro/micro
```

Or via Docker

```shell
docker pull udian/micro
```


> Etcd (服务发现，配置中心)

> [Nsq](http://wiki.jikexueyuan.com/project/nsq-guide/) （消息队列，可以替换成kafka）


### Run Etcd
```
$ docker pull quay.io/coreos/etcd //下载etcd docker镜像
  //docker 启动etcd
  docker run \    
    -d  \
    -p 2379:2379 \
    -p 2380:2380 \
    --name etcd \
    --volume=/tmp/etcd-data.tmp:/etcd-data \
    --restart=always \
    quay.io/coreos/etcd \
    /usr/local/bin/etcd \
    --name etcd0 \
    --data-dir /etcd-data \
    --listen-client-urls http://0.0.0.0:2379 \
    --advertise-client-urls http://0.0.0.0:2379 \
    --listen-peer-urls http://0.0.0.0:2380 \
    --initial-advertise-peer-urls http://0.0.0.0:2380 \
    --initial-cluster etcd0=http://0.0.0.0:2380 \
    --initial-cluster-token tkn
```
### Run nsq
```
$ 
1.下载nsq docker镜像
    docker pull nsqio/nsq 
2.启动nsqlookupd
    docker run -d --link nsqd --name lookupd -p 4160:4160 -p 4161:4161 --restart=always nsqio/nsq /nsqlookupd  
3.启动nsqd
    docker run -d --link lookupd --name nsqd -p 4150:4150 -p 4151:4151 --restart=always nsqio/nsq /nsqd --lookupd-tcp-address=lookupd:4160
4.启动nsqadmin
    docker run -d --link lookupd --link nsqd --name nsqadmin -p 4171:4171 --restart=always nsqio/nsq /nsqadmin --lookupd-http-address=lookupd:4161
```

### Run Service

```
$ go run main.go --registry=etcd --broker=nsq
```

### Building a container

If you would like to build the docker container do the following
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i
docker build -t udian/goods .

```
run in docker 
```
docker run --link etcd --link nsqd -p 50051:50051 --name=goods -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=etcd -e MICRO_REGISTRY_ADDRESS=http://etcd:2379 -e MICRO_BROKER=nsq -e MICRO_BROKER_ADDRESS=nsqd:4150  udian/goods
```

### Run the micro api

启动api服务，并转发到rpc服务
```

micro --registry=etcd  api --handler=rpc

```
Or via Docker

```shell
docker run --link=etcd -p 8080:8080 -e MICRO_REGISTRY=etcd -e MICRO_REGISTRY_ADDRESS=http://etcd:2379 ttouch/micro api --handle=rpc
```

api调用
```
curl -H 'Content-Type: application/json' -d '{"name": "John2","barcoe":"2222"}' http://127.0.0.1:8080/v1/goods/detail

```

