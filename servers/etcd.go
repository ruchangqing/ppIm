package servers

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"ppIm/utils"
	"strings"
)

var EtcdClient *clientv3.Client

//注册本机地址到etcd服务
func RegisterServer() {
	serverIp := utils.GetIntranetIp()
	serverAddress := serverIp + ":" + viper.GetString("cluster.rpc_port")
	AddServer(serverAddress)
	//新建租约
	resp, err := EtcdClient.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	//授予租约
	key := "server_" + serverAddress
	_, err = EtcdClient.Put(context.TODO(), key, "1", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
	//keep-alive
	ch, kaerr := EtcdClient.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	go func() {
		for {
			select {
			case <-ch:
				//keep-alive
			}
		}
	}()
}

//获取所有注册到哦哦哦哦哦etcd的服务器
func WatchServers() {
	rch := EtcdClient.Watch(context.Background(), "server_", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			arr := strings.Split(string(ev.Kv.Key), "server_")
			serverAddr := arr[1]
			switch ev.Type {
			case 0: //put
				AddServer(serverAddr)
				fmt.Println("cluster join:" + serverAddr)
			case 1: //delete
				DelServer(serverAddr)
				fmt.Println("cluster leave:" + serverAddr)
			}
		}
	}
}
