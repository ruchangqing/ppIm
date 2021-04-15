package servers

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"ppIm/utils"
	"time"
)

var EtcdClient *clientv3.Client

//注册本机地址到etcd服务
func RegisterServer() {
	serverIp := utils.GetIntranetIp()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	_, err := EtcdClient.Put(ctx, "server_"+serverIp, "1")
	cancel()
	if err != nil {
		log.Fatal("服务注册到etcd出错：" + err.Error())
	}
}

//获取所有注册到哦哦哦哦哦etcd的服务器
func GetServers() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	resp, err := EtcdClient.Get(ctx, "server_", clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("当前集群：")
	for _, v := range resp.Kvs {
		fmt.Println(string(v.Key))
	}
}
