package servers

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"ppIm/global"
	"strings"
	"time"
)

var EtcdClient *clientv3.Client

//获取当前所有集群
func GetAllServers() []string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	resp, err := EtcdClient.Get(ctx, "server_", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println(err)
	}
	var servers []string
	for _, ev := range resp.Kvs {
		arr := strings.Split(string(ev.Key), "server_")
		serverAddr := arr[1]
		servers = append(servers, serverAddr)
	}
	return servers
}

//注册集群
func RegisterServer() {
	AddServer(global.ServerAddress)
	//新建租约
	resp, err := EtcdClient.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println(err)
	}
	//授予租约
	key := "server_" + global.ServerAddress
	_, err = EtcdClient.Put(context.TODO(), key, "1", clientv3.WithLease(resp.ID))
	if err != nil {
		fmt.Println(err)
	}
	//keep-alive
	ch, kaerr := EtcdClient.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		fmt.Println(kaerr)
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

//发现集群
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
