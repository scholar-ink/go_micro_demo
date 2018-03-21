build:
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i
    docker build -t udian/goods .
run:
    docker run --link etcd --link nsqd -p 50051:50051 --name=goods -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=etcd -e MICRO_REGISTRY_ADDRESS=http://etcd:2379 -e MICRO_BROKER=nsq -e MICRO_BROKER_ADDRESS=nsqd:4150  udian/goods
