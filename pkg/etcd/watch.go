package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

var client *clientv3.Client

func Init() {
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.10.10:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("init etcd client err: %v", err)
	}
}

func WatchKey(key string) {
	wch := client.Watch(context.Background(), key)
	log.Printf("watch %s\n", key)
	for resp := range wch {
		for _, ev := range resp.Events {
			log.Printf("watch, Type %s,Key %s, Value %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
