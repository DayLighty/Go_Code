package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Println("clientv3.New err:", err)
		return
	}
	defer cli.Close()
	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "s4", "真好")
	if err != nil {
		fmt.Println("put err:", err)
		return
	}
	cancel()
	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gr, err := cli.Get(ctx, "s4")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	for _, ev := range gr.Kvs {
		fmt.Printf("key:%s value:%s\n", ev.Key, ev.Value)
	}
	cancel()
}
