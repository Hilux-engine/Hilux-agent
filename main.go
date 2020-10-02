package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	// 创建 api v3 的 client
	cli, err := clientv3.New(clientv3.Config{
		// etcd https api 端点
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	//fmt.Printf("%v", cli)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	rch := cli.Watch(context.Background(), "test")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
